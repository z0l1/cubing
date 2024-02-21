package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n') // reads input until the first occurrence of '\n'
	return strings.TrimRight(text, "\n")
}
