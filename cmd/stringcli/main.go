package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"petgo/internal/stringapp"
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func main() {
	var (
		input      string
		daemonMode bool
	)

	flag.StringVar(&input, "input", "", "input string")
	flag.BoolVar(&daemonMode, "daemon", false, "run in daemon mode")
	flag.Parse()

	if daemonMode {
		runDaemon()
		return
	}

	if input == "" {
		exitWithError(fmt.Errorf("pass a string via --input or run --daemon"))
	}

	result, err := stringapp.Unpack(input)
	if err != nil {
		exitWithError(err)
	}

	fmt.Println(result)
}

func runDaemon() {
	fmt.Println("Ctrl+C to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a string: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				exitWithError(err)
			}
			return
		}

		result, err := stringapp.Unpack(scanner.Text())
		if err != nil {
			logger.Error("failed to unpack input", "err", err)
			continue
		}

		fmt.Println(result)
	}
}

func exitWithError(err error) {
	logger.Error("command failed", "err", err)
	os.Exit(1)
}
