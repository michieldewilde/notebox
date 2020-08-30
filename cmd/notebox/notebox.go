package main

import (
	"flag"
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
	addrFlag := flag.String("http", "0.0.0.0:80", "HTTP address notebox will be exposed on")
	flag.Parse()

	if err := run(os.Stdout, *addrFlag); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(stdout io.Writer, addr string) error {
	s := &server{
		router: mux.NewRouter(),
	}

	s.registerSchema()
	s.registerRoutes()

	return s.listenAndServe(addr)
}

func (s *server) registerRoutes() {
	s.router.HandleFunc("/api", s.handleAPI()).Methods("POST")
	s.router.Handle("/graphiql", s.handleGraphiql("/api"))
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

func (s *server) handleGraphiql(r string) http.Handler {
	return graphqlbackend.HandleGraphiql(r)
}

func (s *server) listenAndServe(addr string) error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on: %v\n", addr)

	return srv.ListenAndServe()
}
