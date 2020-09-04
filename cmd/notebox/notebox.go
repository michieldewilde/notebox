package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	exitFail = 1
)

var (
	addr       = flag.String("http", "0.0.0.0:80", "HTTP address notebox will be exposed on")
	datasource = flag.String("datasource", "", "Connection URL of the datasource")
)

func main() {
	flag.Parse()

	s := &Server{}

	if err := s.run(os.Stdout, *addr, *datasource); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
