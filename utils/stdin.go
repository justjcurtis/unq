/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetStdIn() string {
	var input string
	in := bufio.NewScanner(os.Stdin)
	stats, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if stats.Size() == 0 {
		os.Exit(0)
	}
	i := 0
	for in.Scan() {
		text := strings.TrimSpace(in.Text())
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
