package main

import (
	"errors"
	"fmt"
	"os"

	"./indent"
	"./report"
)

const (
	Usage = "Usage: %s command args...\n"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0])
		fmt.Fprintf(os.Stderr, indent.Usage(), os.Args[0])
		fmt.Fprintf(os.Stderr, report.Usage(), os.Args[0])
		return 1
	}

	command := os.Args[1]

	var e error
	switch command {

	case "indent":
		e = indent.Process(os.Args[2:])
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			fmt.Fprintf(os.Stderr, indent.Usage(), os.Args[0])
			return 2
		}

	case "report":
		e = report.Process(os.Args[2:])
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			fmt.Fprintf(os.Stderr, report.Usage(), os.Args[0])
			return 2
		}

	default:
		e = errors.New("unknown command")
	}

	return 0
}
