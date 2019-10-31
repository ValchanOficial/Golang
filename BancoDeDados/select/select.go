package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", 3)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
