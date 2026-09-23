package main

import (
	"database/sql"
	"fmt"
	"time"
)

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
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TodoCreateRequest 创建任务请求
type TodoCreateRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
	Priority    string `json:"priority" binding:"oneof=low medium high"`
	DueDate     string `json:"due_date"`
}

// TodoUpdateRequest 更新任务请求
type TodoUpdateRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	State       *string `json:"state,omitempty" binding:"oneof=pending completed cancelled"`
	Priority    *string `json:"priority,omitempty" binding:"oneof=low medium high"`
	DueDate     *string `json:"due_date,omitempty"`
}

// APIResponse 通用API响应
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// PaginatedResponse 分页响应
type PaginatedResponse struct {
	Items     interface{} `json:"items"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
}

// 创建数据库表
func createTable(db *sql.DB) error {
	query := `
	create table if not exists todos(
	id integer primary key autoincrement,
	title varchar(200) not null,
	description text,
	status varchar(20) default 'pending' check(status in('pending','completed','cancelled')),
	priority charchar(10) default 'medium' check(priotity in('low','medium','high')),
	due_date date,
	create_at datetime default current_timesatmp,
	updated_at datetime default current_timesatmp,
	complete_at datetime
	);
`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("创建任务表失败:%w", err)
	}

}
