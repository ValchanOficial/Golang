package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Criando Schema e Tabela")
	db, err := sql.Open("mysql", "root:root@/")
	checkErr(err)
	defer db.Close()
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios(
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`)
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	checkErr(err)
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
