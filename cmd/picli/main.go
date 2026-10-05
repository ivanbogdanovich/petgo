package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"petgo/internal/piapp"
)

func main() {
	workers := flag.Int(
		"workers",
		10,
		"count of workers",
	)

	fmt.Println("workers:", *workers)

	precision := flag.Int(
		"precision",
		7,
		"count of decimal places",
	)

	timeout := flag.Duration(
		"timeout",
		200,
		"time of work",
	)
	fmt.Println("timeout", timeout)

	verbose := flag.Bool(
		"verbose",
		false,
		"print count of iterations",
	)

	flag.Parse()

	if *workers < 1 {
		fmt.Fprintln(os.Stderr, "workers must be greater than 0")
		os.Exit(2)
	}

	if *precision < 0 {
		fmt.Fprintln(os.Stderr, "precision not be negative")
		os.Exit(2)
	}

	if *timeout < 0 {
		fmt.Fprintln(os.Stderr, "timeout not be negative")
		os.Exit(2)
	}

	ctx := context.Background()
	cancel := func() {}

	if *timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, *timeout)
	}
	defer cancel()

	ctx, stopSignals := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stopSignals()

	startedAt := time.Now()

	result := piapp.Run(ctx, *workers)

	elapsed := time.Since(startedAt)

	fmt.Printf("PI %.10f\n", result.Pi)

	if *verbose {
		fmt.Printf("iterations: %d\n", result.Iterations)
		fmt.Printf("elapsed: %s\n", elapsed)
	}
}
