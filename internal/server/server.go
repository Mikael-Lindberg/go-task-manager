package server

import (
	"net/http"

	"github.com/Mikael-Lindberg/task-manager-api/internal/handlers"
	"github.com/gorilla/mux"
)

// NewRouter creates a new HTTP router and configures routes.
func NewRouter(taskHandler *handlers.TaskHandler) http.Handler {
	r := mux.NewRouter()

	// Define routes and attach handlers
	r.HandleFunc("/api/tasks", taskHandler.CreateTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks/{id}", taskHandler.GetTaskByIDHandler).Methods(http.MethodGet)
	// Add routes for updating and deleting tasks

	return r
}
