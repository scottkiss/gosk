package main

import (
	"flag"
	"fmt"
	"github.com/scottkiss/gosk"
	"os"
)

const VERSION = "0.0.1"

const (
	USAGE = `
gosk is a static site generator in Go

Usage:

        gosk command [args...]

The commands are:

	build	        			build and generate site.
	run[address:port]  	        e.g. gosk run :80 (default listen at port 8080)
	version         			print gosk version

`
)

var httpAddr = ":8080"

func main() {
	flag.Parse()
	args := flag.Args()
	argsLength := len(args)
	if argsLength == 0 || argsLength > 3 {
		Usage()
		os.Exit(1)
	}
	switch args[0] {
	default:
		Usage()
		os.Exit(1)
	case "build":
		gosk.Build()
	case "run":
		if args[1] != "" {
			httpAddr = args[1]
		}
		fmt.Println("Listen at ", httpAddr)
		gosk.Run(httpAddr)
	case "version":
		fmt.Print("gosk version " + VERSION)
	}
}

func Usage() {
	fmt.Println(USAGE)
}
