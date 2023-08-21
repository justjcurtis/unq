/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import "strings"

func CleanUp(lines []string) []string {
	var cleaned []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		trimmed := strings.TrimSpace(line)
		cleaned = append(cleaned, trimmed)
	}
	return cleaned
}
