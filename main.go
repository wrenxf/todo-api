package main

import "time"

// Todo 任务结构体
type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" binding:"required,min=1,max=200"`
	Description string     `json:"description" binding:"max=1000"`
	State       string     `json:"state" binding:"oneof=pending completed cancelled"`
	Priority    string     `json:"priority" binding:"oneof=low medium high"`
	DueDate     *string    `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdateAt    time.Time  `json:"update_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
