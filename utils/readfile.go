/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"os"
	"strings"
)

func GetFileContents(file string, trim bool) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	if len(contents) == 0 {
		return "", nil
	}
	strContents := string(contents[:len(contents)-1])
	if trim {
		lines := strings.Split(strContents, "\n")
		for i, line := range lines {
			lines[i] = strings.TrimSpace(line)
		}
		strContents = strings.Join(lines, "\n")
	}
	return strContents, nil
}
