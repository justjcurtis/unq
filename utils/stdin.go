/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetStdIn(trim bool) string {
	var input string
	in := bufio.NewScanner(os.Stdin)
	stats, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if stats.Size() == 0 {
		return ""
	}
	i := 0
	for in.Scan() {
		text := in.Text()
		if trim {
			text = strings.TrimSpace(text)
		}
		if text == "" || text == "\n" {
			continue
		}
		if i > 0 {
			input += "\n"
		}
		input += text
		i++
	}
	return input
}
