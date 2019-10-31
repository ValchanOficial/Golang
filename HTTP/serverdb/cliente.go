package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//Usuario : usuario
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

//UsuarioHandler : analise o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		find(w, r, id)
	case r.Method == "GET":
		findAll(w, r)
	case r.Method == "POST":
		create(w, r)
	case r.Method == "PUT" && id > 0:
		update(w, r, id)
	case r.Method == "DELETE" && id > 0:
		delete(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Desculpa... :(")
	}
}

func checkIfUserExists(id int) bool {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)
	if isEmpty(u) {
		return true
	} else {
		return false
	}
}

func find(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)
	if isEmpty(u) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Não existe usuário com este id: ", id)
	} else {
		json, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(json))
	}
}

func findAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	checkErr(err)
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Preencha com os dados!")
	} else {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		checkErr(err)

		var usuario Usuario
		err = json.Unmarshal(b, &usuario)
		checkErr(err)

		db, err := sql.Open("mysql", "root:root@/cursogo")
		checkErr(err)
		defer db.Close()

		stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
		res, _ := stmt.Exec(usuario.Nome)
		id, _ := res.LastInsertId()

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "Usuario inserido com sucesso: ", id)
	}
}

func update(w http.ResponseWriter, r *http.Request, id int) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Preencha com os dados!")
	} else {
		if checkIfUserExists(id) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Não existe usuário com este id: ", id)
		} else {
			b, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			checkErr(err)

			var usuario Usuario
			err = json.Unmarshal(b, &usuario)
			checkErr(err)

			db, err := sql.Open("mysql", "root:root@/cursogo")
			checkErr(err)
			defer db.Close()

			stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
			stmt.Exec(usuario.Nome, id)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, "Usuario atualizado com sucesso!")
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request, id int) {
	if checkIfUserExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Não existe usuário com este id: ", id)
	} else {
		db, err := sql.Open("mysql", "root:root@/cursogo")
		checkErr(err)
		defer db.Close()

		stmt, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
		stmt.Exec(id)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "Usuario deletado com sucesso!")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isEmpty(usuario Usuario) bool {
	if (Usuario{}) == usuario {
		return true
	}
	return false
}
