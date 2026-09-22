## 🎯 项目概述

### 项目功能

本项目实现一个完整的任务管理 API，包括：

- ✅ **任务管理** - 创建、读取、更新、删除任务
- ✅ **状态管理** - pending、completed、cancelled 三种状态
- ✅ **优先级管理** - low、medium、high 三个优先级
- ✅ **截止日期** - 支持任务截止日期设置
- ✅ **搜索功能** - 按标题和描述搜索任务
- ✅ **分页查询** - 支持大数据量的分页显示
- ✅ **过滤功能** - 按状态、优先级过滤
- ✅ **统计功能** - 任务统计数据

### 技术栈

- **Web 框架**: Gin
- **数据库**: SQLite
- **ORM**: database/sql（标准库）
- **中间件**: CORS、日志、错误处理
- **数据验证**: Gin 绑定验证

## 🏗️ 项目结构



```
01-todo-api/
├── main.go           # 主程序文件
├── go.mod           # Go 模块文件
├── README.md        # 项目文档
├── test-api.sh      # API 测试脚本
└── todos.db         # SQLite 数据库文件（运行时生成）
```



### 代码架构



```
main.go
├── 数据模型 (Models)
│   ├── Todo              # 任务结构体
│   ├── TodoCreateRequest # 创建请求结构体
│   ├── TodoUpdateRequest # 更新请求结构体
│   ├── APIResponse      # API响应结构体
│   └── PaginatedResponse # 分页响应结构体
├── 服务层 (Services)
│   ├── TodoService       # 任务服务接口
│   └── TodoServiceImpl  # 任务服务实现
├── 路由层 (Handlers)
│   ├── handleHome        # 首页处理
│   ├── handleHealth      # 健康检查
│   ├── handleListTodos   # 获取任务列表
│   ├── handleCreateTodo  # 创建任务
│   ├── handleGetTodo     # 获取任务详情
│   ├── handleUpdateTodo  # 更新任务
│   ├── handleDeleteTodo  # 删除任务
│   ├── handleToggleTodo  # 切换任务状态
│   └── handleTodoStatistics # 统计信息
└── 中间件 (Middleware)
    ├── corsMiddleware    # CORS 跨域中间件
    ├── loggingMiddleware # 日志中间件
    └── errorHandlerMiddleware # 错误处理中间件
```