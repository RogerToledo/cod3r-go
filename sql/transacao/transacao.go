package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nomes := []string{"Cleide", "Daniela","Ednei"}
	nomesMap := criaMap(2004, nomes)

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios values(?,?)")

	for id, nome := range nomesMap {
		_, err := stmt.Exec(id, nome)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	_, errExec := stmt.Exec(1, "Liliane") // Duplicado para testar rollback
	if errExec != nil {
		tx.Rollback()
		log.Fatal(errExec)
	}

	tx.Commit()
}

func criaMap(init int, nomes []string) map[int]string {
	i := init
	var nomesMap = make(map[int]string)
	for _, nome := range nomes {
		nomesMap[i] = nome
		i++
	}

	return nomesMap
}