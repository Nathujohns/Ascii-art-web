package asciiart

import (
	"bufio"
	"os"
	"strings"
)

// GetAsciiLine reads a specific line number from a file, where line numbering starts from 0.
func GetAsciiLine(filename string, num int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		if lineNum == num {
			return scanner.Text(), nil
		}
		lineNum++
	}
	return "", os.ErrNotExist // or a custom error indicating the line number was not found
}

// AsciiArt generates ASCII art from the input text using a specified font stored in a file.
func AsciiArt(input, fontName string) (string, error) {
	banner := "banners/" + fontName + ".txt"
	var result strings.Builder

	lines := strings.Split(input, "\n")
	for _, word := range lines {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				asciiLine, err := GetAsciiLine(banner, 1+int(letter-' ')*9+i)
				if err != nil {
					return "", err
				}
				result.WriteString(asciiLine)
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
