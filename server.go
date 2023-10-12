package main

import (
	"fmt"
	"go-gql/todox"
	"go-gql/todox/resolvers"
	"go-gql/todox/store"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(todox.NewExecutableSchema(todox.Config{
		Resolvers: &resolvers.Resolver{Store: store.New()},
	}))

	router := chi.NewRouter()
	router.Use(auth())

	router.Handle("/", playground.Handler("TodoX Playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("auth-token")
			fmt.Println("AUTH_TOKEN in MW: ", token)

			if len(token) == 0 || token != "secrettoken" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
