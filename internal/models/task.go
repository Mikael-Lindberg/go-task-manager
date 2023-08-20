package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(title, content string) *Task {
	return &Task {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}