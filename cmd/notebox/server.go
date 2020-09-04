package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/michieldewilde/notebox/pkg/graphqlbackend"
)

type Server struct {
	router *mux.Router
	schema *graphql.Schema
	db     string
}

func (s *Server) run(stdout io.Writer, addr string, db string) error {
	s.router = mux.NewRouter()

	s.registerSchema()
	s.registerRoutes()

	return s.listenAndServe(addr)
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/api", s.handleAPI()).Methods("POST")
	s.router.Handle("/graphiql", s.handleGraphiql("/api"))
}

func (s *Server) registerSchema() error {
	schema, err := graphqlbackend.NewSchema()

	if err != nil {
		return err
	}

	s.schema = schema
	return nil
}

func (s *Server) handleAPI() http.HandlerFunc {
	return graphqlbackend.Handle(s.schema)
}

func (s *Server) handleGraphiql(r string) http.Handler {
	return graphqlbackend.HandleGraphiql(r)
}

func (s *Server) listenAndServe(addr string) error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on: %v\n", addr)

	return srv.ListenAndServe()
}
