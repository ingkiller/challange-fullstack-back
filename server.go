package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/ingkiller/hackernews/graph"
	"github.com/ingkiller/hackernews/graph/generated"
	"github.com/ingkiller/hackernews/internal/auth"
	"github.com/rs/cors"
	"net/http"
)

const defaultPort = ":8080"

func main() {
	router := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,

		Debug: true,
	}).Handler)

	router.Use(auth.Middleware())
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost:8080"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL Playground", "/api"))
	router.Handle("/api", srv)

	err := http.ListenAndServe(defaultPort, router)
	if err != nil {
		panic(err)
	}
}
