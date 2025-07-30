# 📝 Todo API in Go (Golang)

A simple RESTful Todo API built with **Go**, **Gin**, and **MySQL**.

This project is meant for learning and can be extended with authentication (JWT), pagination, and more.

---

## 📦 Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) – HTTP web framework
- [GORM](https://gorm.io/) – ORM for Go
- [MySQL](https://www.mysql.com/)
- [godotenv](https://github.com/joho/godotenv) – .env config loader

---

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/awandha/my-todo-api-go
cd my-todo-api-go
```

### 2. Create a .env file in the project root
```DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=todo_db
```
Replace your_password with your MySQL root password.

### 3. Create the MySQL database
```
CREATE DATABASE todo_db;
```
You can use MySQL CLI or a GUI like MySQL Workbench/phpMyAdmin.

### 4. Install Go dependencies
```
go mod tidy
```

### 5. Run the application
```
go run main.go
```
The server will start at:
👉 http://localhost:8080

📚 API Endpoints
➕ Create a Todo
POST /todos/

Request Body:
```
{
  "title": "Learn Go",
  "completed": false
}
```

📥 Get All Todos
GET /todos/

🔍 Get Todo by ID
GET /todos/:id

🖊️ Update a Todo
PUT /todos/:id

Request Body:
```
{
  "title": "Updated Title",
  "completed": true
}
```

❌ Delete a Todo
DELETE /todos/:id

🛡️ License
MIT – Free to use and modify.

🙋‍♂️ Author
Made with ❤️ by Awandha