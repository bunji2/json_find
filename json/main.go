package main

import (
	"errors"
	"fmt"
	"os"

	"./indent"
	"./path"
	"./report"
)

const (
	Usage = "Usage: %s command args...\n"
)

var conf Config

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0])
		fmt.Fprintf(os.Stderr, indent.Usage(), os.Args[0])
		fmt.Fprintf(os.Stderr, report.Usage(), os.Args[0])
		fmt.Fprintf(os.Stderr, path.Usage(), os.Args[0])
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

	case "path":
		e = path.Process(os.Args[2:])
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			fmt.Fprintf(os.Stderr, path.Usage(), os.Args[0])
			return 2
		}
	case "server":
		e = Load_config(&conf)
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			return 2
		}
		e = Process_server()
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			fmt.Fprintf(os.Stderr, Usage_server, os.Args[0])
			return 2
		}

	default:
		e = errors.New("unknown command")
	}

	return 0
}
