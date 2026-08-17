package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/danielronalds/danielronalds.github.io/internal/parsing"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: blog <src-dir> <dest-dir>")
		os.Exit(errCodeInvalidArguments)
	}

	srcDir := os.Args[1]
	destDir := os.Args[2]

	routes, err := parsing.ParseRoutes(slog.Default(), srcDir, destDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err.Error())
		os.Exit(errCodeFailedToParseRoutes)
	}

	for _, route := range routes {
		fmt.Printf("%s -> %s\n", route.Src, route.Dest)
	}
}
