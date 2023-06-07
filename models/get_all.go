package models

import (
	"log"

	"github.com/Jonathansoufer/go-postgres-api/db"
)

func GetAll()(todos []Todo, err error){
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM todos")

	if err != nil {
		return
	}

	for rows.Next(){
		var t Todo
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed)

		if err != nil {
			log.Default().Println(err)
			continue
		}

		todos = append(todos, t)
	}

	return todos, err
}