/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

func GetUnique(lines []string) ([]string, map[string]int) {
	var unique []string
	m := make(map[string]int)
	for _, line := range lines {
		if _, ok := m[line]; ok {
			m[line]++
			continue
		}
		m[line] = 1
		unique = append(unique, line)
	}
	return unique, m
}
