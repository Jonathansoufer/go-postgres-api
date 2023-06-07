package models

import "github.com/Jonathansoufer/go-postgres-api/db"

func Delete(id int64)(int64, error){
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "DELETE FROM todos WHERE id = $1"

	result, err := conn.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}