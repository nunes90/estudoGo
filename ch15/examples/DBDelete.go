package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	deleteStatement := `
	DELETE FROM test
	WHERE id = $1
	`
	deleteResult, deleteResultErr := db.Exec(deleteStatement, 2)
	if deleteResultErr != nil {
		panic(deleteResultErr)
	}
	deletedRecords, deletedRecordsErr := deleteResult.RowsAffected()
	if deletedRecordsErr != nil {
		panic(deletedRecordsErr)
	}
	fmt.Println("Number of records deleted: ", deletedRecords)
	defer db.Close()
}
