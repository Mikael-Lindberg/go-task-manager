package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Mikael-Lindberg/task-manager-api/internal/handlers"
	"github.com/Mikael-Lindberg/task-manager-api/internal/repositories"
	"github.com/Mikael-Lindberg/task-manager-api/internal/server"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Database setup
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Repository setup
	taskRepo := repositories.NewTaskRepository(db)

	// Handler setup
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// HTTP server setup
	router := server.NewRouter(taskHandler)
	serverAddr := "localhost:8080"

	fmt.Printf("Server listening on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
