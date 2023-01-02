package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id int
	nome string
}

func main() {
	db, errDb := sql.Open("mysql", "root:1234@/cursogo")
	if errDb != nil {
		log.Fatal()
	}

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 20)
	defer rows.Close()

	var usuarios []usuario

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		usuarios = append(usuarios, u)
	}

	fmt.Println(usuarios)

	for _, usr := range usuarios {
		fmt.Printf("Indice: %v, Nome: %v\n", usr.id, usr.nome)
	}
}