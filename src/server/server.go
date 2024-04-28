package server

import (
	"html/template"
	"net/http"
	"strings"

	"ascii/src/asciiart"
)

// ErrorInfo holds details about an error to be displayed
type ErrorInfo struct {
	StatusCode string
	Message    string
}

// ASCIIResult holds data for rendering ASCII art result
type ASCIIResult struct {
	Text  string
	Style string
	Art   string
}

// renderErrorPage renders an HTML page for errors
func renderErrorPage(w http.ResponseWriter, r *http.Request, errorInfo *ErrorInfo) {
	tmpl, _ := template.ParseFiles("templates/error.html")
	tmpl.Execute(w, errorInfo)
}

// HomePageHandler handles the root path and renders the main page
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		httpError := ErrorInfo{StatusCode: "404", Message: "Page not found"}
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, r, &httpError)
		return
	}
	if r.Method != "GET" {
		httpError := ErrorInfo{StatusCode: "405", Message: "Method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		renderErrorPage(w, r, &httpError)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		httpError := ErrorInfo{StatusCode: "500", Message: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, r, &httpError)
		return
	}
	tmpl.Execute(w, nil)
}

// ASCIIArtHandler processes the ASCII art creation form
func ASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		httpError := ErrorInfo{StatusCode: "500", Message: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, r, &httpError)
		return
	}
	inputText := r.PostFormValue("input-text")
	// Normalizing the input by removing new line characters
	normalizedInput := strings.ReplaceAll(inputText, "\r\n", "")
	// Character validation
	for _, char := range normalizedInput {
		if char < 32 || char > 126 {
			httpError := ErrorInfo{StatusCode: "400", Message: "Invalid input"}
			w.WriteHeader(http.StatusBadRequest)
			renderErrorPage(w, r, &httpError)
			return
		}
	}
	style := r.PostFormValue("banner")
	if style != "standard" && style != "shadow" && style != "thinkertoy" {
		httpError := ErrorInfo{StatusCode: "404", Message: "Banner style not found"}
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, r, &httpError)
		return
	}
	asciiArt, artErr := asciiart.AsciiArt(inputText, style)
	if artErr != nil {
		httpError := ErrorInfo{StatusCode: "500", Message: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, r, &httpError)
		return
	}
	resultTemplate := template.Must(template.ParseFiles("templates/ascii-art.html"))
	resultData := ASCIIResult{Text: inputText, Style: style, Art: asciiArt}
	resultTemplate.Execute(w, resultData)
}
