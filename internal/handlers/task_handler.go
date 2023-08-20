package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mikael-Lindberg/task-manager-api/internal/models"
	"github.com/Mikael-Lindberg/task-manager-api/internal/repositories"
)

type TaskHandler struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskHandler(taskRepo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepo}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.taskRepo.CreateTask(&task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := h.taskRepo.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for updating a task here
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for deleting a task here
}