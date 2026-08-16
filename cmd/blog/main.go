package main

import (
	"fmt"
	"os"

	"github.com/danielronalds/danielronalds.github.io/internal/parsing"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "no src dir provided")
		os.Exit(errCodeNoSrcDir)
	}

	srcDir := os.Args[1]

	routes, err := parsing.ParseRoutes(srcDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error occured: %s\n", err.Error())
		os.Exit(errCodeFailedToParseRoutes)
	}

	for _, route := range routes {
		fmt.Println("%s -> %s", route.Src, route.Dest)
	}
}
