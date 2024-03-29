package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	book_api "github.com/dominik-robert/book-api"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("Graphql playground", "/query"))
	http.Handle("/query", handler.GraphQL(book_api.NewExecutableSchema(book_api.Config{Resolvers: &book_api.Resolver{}})))
	//http.Handle("/query", handler.GraphQL(book_api.NewExecutableSchema(book_api.Config{Resolvers: &book_api.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
