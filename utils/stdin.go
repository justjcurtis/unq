/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"bufio"
	"os"
)

func GetStdIn() []string {
	// TODO: check if nothing is piped in && exit
	var lines []string
	in := bufio.NewScanner(os.Stdin)
	stats, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if stats.Size() == 0 {
		os.Exit(0)
	}
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return CleanUp(lines)
}
