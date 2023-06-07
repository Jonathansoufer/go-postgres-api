package models

import "github.com/Jonathansoufer/go-postgres-api/db"

func Get(id int64)(todo Todo, err error){
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)

	return todo, err
}