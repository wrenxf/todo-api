package main

import (
	"database/sql"
	"fmt"
	"log"
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
	return nil
}

// 初始化数据库
func initDatabase() *sql.DB {
	//打开数据库连接
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	//检查连接是否健康,测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}

	//建表
	if err := createTable(db); err != nil {
		log.Fatal("数据库创建失败:", err)
	}

	return db
}

// TodoService 任务服务接口
type TodoService interface {
	Create(todo *Todo) error
	GetByID(id int) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
	List(filter TodoFilter) ([]Todo, int, error)
	ToggleStatus(id int, status string) error
}

// TodoFilter 任务查询过滤器
type TodoFilter struct {
	Status   string
	Priority string
	Search   string
	Page     int
	PageSize int
}

// TodoServiceImpl 任务服务实现
// TodoServiceImpl 是任务服务接口的具体实现结构体，内部持有数据库连接
type TodoServiceImpl struct {
	db *sql.DB
}

// NewTodoService 是任务服务的构造函数，通过依赖注入的方式接收数据库连接，并返回服务接口
func NewTodoService(db *sql.DB) TodoService {
	return &TodoServiceImpl{db: db}
}

func (s *TodoServiceImpl) Create(todo *Todo) error {
	query := `
	insert into todos(title,description,status,priority,due_date,created_at,updated_at)
	values(?,?,?,?,?,?,?)
`
	now := time.Now()
	todo.CreatedAt = now
	todo.UpdateAt = now
	if todo.State == "" {
		todo.State = "pending"
	}
	result, err := s.db.Exec(query, todo.Title, todo.Description, todo.State, todo.Priority, todo.DueDate, todo.CreatedAt, todo.UpdateAt)
	if err != nil {
		return fmt.Errorf("创建任务失败: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取任务ID失败: %w", err)
	}
	todo.ID = int(id)
	return nil
}

func (t TodoServiceImpl) GetByID(id int) (*Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoServiceImpl) Update(todo *Todo) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoServiceImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoServiceImpl) List(filter TodoFilter) ([]Todo, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoServiceImpl) ToggleStatus(id int, status string) error {
	//TODO implement me
	panic("implement me")
}
