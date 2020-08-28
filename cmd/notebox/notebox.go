package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/michieldewilde/notebox/pkg/http/graphql"
)

const (
	exitFail = 1
)

type server struct {
	router *mux.Router
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(stdout io.Writer) error {
	s := &server{
		router: mux.NewRouter(),
	}

	s.registerRoutes()

	return http.ListenAndServe("localhost:8080", s.router)
}

func (s *server) registerRoutes() {
	s.router.HandleFunc("/api", s.handleAPI())
}

func (s *server) handleAPI() http.HandlerFunc {
	return graphql.Handle()
}
