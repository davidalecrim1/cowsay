package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/davidalecrim1/cowsay/cowsay"
)

func main() {
	// this checks if the program is being run with a pipe
	info, err := os.Stdin.Stat()
	if err != nil {
		slog.Error("failed to get stdin info", "error", err)
		os.Exit(1)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'Hello, world!' | cowsay")
		os.Exit(1)
	}

	output := []string{}
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		slog.Error("failed to read stdin", "error", err)
		os.Exit(1)
	}

	generated := cowsay.GenerateCowSay(output)
	fmt.Print(generated)
}
