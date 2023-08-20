package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Mikael-Lindberg/task-manager-api/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	query := `
		INSERT INTO tasks (title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $3)
		RETURNING id
	`

	now := time.Now()
	err := repo.db.QueryRow(query, task.Title, task.Content, now, now).Scan(&task.ID)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return fmt.Errorf("error creating task")
	}

	return nil
}

func (repo *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	task := &models.Task{}
	err := repo.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found")
		}
		log.Printf("Error retrieving task: %v", err)
		return nil, fmt.Errorf("error retrieving task")
	}

	return task, nil
}

func (repo *TaskRepository) UpdateTask(task *models.Task) error {
	query := `
		UPDATE tasks
		SET title = $2, content = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := repo.db.Exec(query, task.ID, task.Title, task.Content, time.Now())
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return fmt.Errorf("error updating task")
	}

	return nil
}

func (repo *TaskRepository) DeleteTask(id int) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := repo.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return fmt.Errorf("error deleting task")
	}

	return nil
}