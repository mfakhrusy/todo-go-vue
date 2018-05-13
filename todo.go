// todo.go
package main

import (
	"database/sql"

	"github.com/mfakhrusy/todo-go-vue/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	// create new instance of Echo
	e := echo.New()

	e.File("/", "public.index.html")
	e.GET("/tasks", handlers.GetTask(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks", handlers.DeleteTask(db))

	e.Start(":1323")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// check any db error then exit
	if err != nil {
		panic(err)
	}

	// check if we  can connect to the db
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	)
	`

	_, err := db.Exec(sql)

	// exit if something goes wrong
	if err != nil {
		panic(err)
	}
}
