package main

import (
	"Sistem-Absensi-Backend-Go/Database"
	"Sistem-Absensi-Backend-Go/Routes"
	"Sistem-Absensi-Backend-Go/graph"
	"Sistem-Absensi-Backend-Go/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	Database.Connect()
	// Inisialisasi router Gin
	router := gin.Default()
	router.Use(cors.Default())
	// Register semua routes
	Routes.SetupRoutes(router)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	// Optional: GraphQL Playground
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	// Playground & GraphQL
	router.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/graphql")))
	router.POST("/graphql", gin.WrapH(srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

//func main() {

//Database.Connect()
//// Inisialisasi router Gin
//router := gin.Default()
//
//// Register semua routes
//Routes.SetupRoutes(router)
//
//// Menjalankan HTTP server
////http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
////	fmt.Fprintln(w, "Sistem Absensi API is running")
////})
//
//fmt.Println("Server started on port 8080")
//err := router.Run(":8080")
//if err != nil {
//	log.Fatal("Server error:", err)
//}

//}
