package main

import (
	"encoding/json"
	"fmt"
	"go-gql/todox"
	"go-gql/todox/resolvers"
	"go-gql/todox/store"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const port = "8080"

func main() {
	srv := handler.NewDefaultServer(todox.NewExecutableSchema(todox.Config{
		Resolvers: &resolvers.Resolver{Store: store.New()},
	}))

	router := chi.NewRouter()
	router.Use(auth())

	router.Handle("/", playground.Handler("TodoX Playground", "/query"))
	router.Handle("/graph", srv)
	router.Post("/raw", RAWhandler)
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Just to investigate what actually comes in the request
func RAWhandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		fmt.Println("NO body")
	}

	m := map[string]any{}
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		fmt.Println("Err in reading body")
		return
	}

	for k, v := range m {
		fmt.Printf("KEY: %s\nVAL: %v", k, v)
		fmt.Printf("Value type: %T", v)
	}
}

// Sample middleware to authenticate the request
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
