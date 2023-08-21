/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"errors"
	"strings"
)

func GetDelimiter(lines []string) (string, error) {
	length := len(lines)

	if length == 0 {
		return "", errors.New("no lines to parse")
	}

	if length == 1 {
		commaCount := 0
		spaceCount := 0
		commaSpaceCount := 0
		lastChar := ""
		for _, char := range lines[0] {
			if char == ',' {
				commaCount++
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
		if commaCount > 0 {
			if commaSpaceCount > 0 {
				return ", ", nil
			}
			return ",", nil
		}
		if spaceCount > 0 {
			return " ", nil
		}
	}

	if length > 1 {
		commaCount := 0
		for _, line := range lines {
			if strings.HasSuffix(line, ",") {
				commaCount++
			}
		}
		if commaCount >= length-1 {
			return ",\n", nil
		}
		return "\n", nil
	}

	return "", errors.New("could not determine delimiter")
}
