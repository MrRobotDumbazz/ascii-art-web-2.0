package service

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	ErrEmptyValue        = errors.New("Empty value")
	ErrNotPartAsciiTable = errors.New("Your [STRING] is not part of ascii-table...")
)

func AsciiPrinter(str string, asciiMap map[rune][]string) string {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	words := strings.Split(str, "\n")
	var res string
	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, ch := range word {
					val, exists := asciiMap[ch]
					if !exists {
						continue
					}
					res += val[i]
				}
				res += "\n"
			}
		} else {
			res += "\n"
		}
	}
	return res
}

func AsciiMapper(file string) (map[rune][]string, error) {
	asciiMap := map[rune][]string{}

	allValues, err := ReadFile(file)
	if err != nil {
		return nil, err
	}

	ch := ' '
	for len(allValues) != 0 {
		values := allValues[1:9]
		asciiMap[ch] = values
		allValues = allValues[9:]
		ch++
	}
	return asciiMap, nil
}

func IsValidInput(inputStr string) error {
	count := 0
	countN := 0

	if inputStr == "" {
		return ErrEmptyValue
	}
	inputStr = strings.ReplaceAll(inputStr, "\r\n", "\\r\\n")

	for _, r := range inputStr {
		if r != '\n' && r != '\r' && (r < ' ' || r > '~') {
			return ErrNotPartAsciiTable
		}
		if r == 'n' || r == '\n' {
			countN++
		}
		if r == '\\' {
			count++
		}
	}
	return nil
}

func ReadFile(file string) ([]string, error) {
	f, err := os.Open("./internal/fonts/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var allValues []string
	for scanner.Scan() {
		allValues = append(allValues, scanner.Text())
	}
	return allValues, nil
}

func RunAscii(s string, fonts string) (string, error) {
	inputStr := s
	if inputStr == "" {
		return "", nil
	}
	err := IsValidInput(inputStr)
	if err != nil {
		return "", err
	}
	banner, err := BannerCheck(fonts)
	if err != nil {
		return "", err
	}
	asciiMap, err := AsciiMapper(banner)
	if err != nil {
		return "", err
	}
	result := AsciiPrinter(inputStr, asciiMap)
	return result, nil
}

func BannerCheck(fileName string) (string, error) {
	err := errors.New("wrong banner type... try standard/shadow/thinkertoy/yourstyle")
	if fileName == "standard" || fileName == "shadow" || fileName == "thinkertoy" || fileName == "yourstyle" {
		switch {
		case fileName == "standard" || fileName == "shadow" || fileName == "thinkertoy" || fileName == "yourstyle":
			return fileName, nil
		default:
			return "standard", nil
		}
	} else {
		return "", err
	}
}
