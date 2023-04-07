package main

import (
	"bufio"
	"fmt"
	"os"
)

func msgInput() string {
	fmt.Println("Input your message below.")
	fmt.Println("To finish input and send the message, type `:` in a new line and press `Enter`.")
	fmt.Println("--------------------Your Message-------------------")

	scanner := bufio.NewScanner(os.Stdin)
	var msg string
	for scanner.Scan() {
		line := scanner.Text()
		if line == ":" {
			break
		}
		msg += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return msg
}
