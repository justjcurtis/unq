/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"errors"
)

func GetDelimiter(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("no input to parse")
	}

	commaCount := 0
	spaceCount := 0
	commaSpaceCount := 0
	newlineCount := 0
	commaNewlineCount := 0
	lastChar := ""
	for _, char := range input {
		if char == ',' {
			commaCount++
		}
		if char == '\n' {
			newlineCount++
			if lastChar == "," {
				commaNewlineCount++
			}
		}
		if char == ' ' {
			if lastChar == "," {
				commaSpaceCount++
			} else {
				spaceCount++
			}
		}
		lastChar = string(char)
	}
	if newlineCount > 0 {
		if commaNewlineCount > 0 {
			return ",\n", nil
		}
		return "\n", nil
	}
	if commaCount > 0 {
		if commaSpaceCount > 0 {
			return ", ", nil
		}
		return ",", nil
	}
	if spaceCount > 0 {
		return " ", nil
	}

	return "", errors.New("could not determine delimiter")
}
