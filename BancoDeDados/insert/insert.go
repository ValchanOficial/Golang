package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria")
	stmt.Exec("Felipe")
	stmt.Exec("João")
	stmt.Exec("Rosa")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
