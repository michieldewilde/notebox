package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/michieldewilde/notebox/pkg/graphqlbackend"
)

const (
	exitFail = 1
)

type server struct {
	router *mux.Router
	schema *graphql.Schema
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

	s.registerSchema()
	s.registerRoutes()

	return s.listenAndServe(8080)
}

func (s *server) registerRoutes() {
	s.router.HandleFunc("/api", s.handleAPI())
}

func (s *server) registerSchema() error {
	schema, err := graphqlbackend.NewSchema()

	if err != nil {
		return err
	}

	s.schema = schema
	return nil
}

func (s *server) handleAPI() http.HandlerFunc {
	return graphqlbackend.Handle(s.schema)
}

func (s *server) listenAndServe(p int) error {
	addr := fmt.Sprintf("127.0.0.1:%d", p)
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on: %v\n", addr)

	return srv.ListenAndServe()
}
