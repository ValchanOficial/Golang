package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Valchan", 1)
	stmt.Exec("Sheristone", 2)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
