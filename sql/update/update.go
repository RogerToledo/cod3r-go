package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, errDb := sql.Open("mysql", "root:1234@/cursogo")
	if errDb != nil {
		log.Fatal()
	}
	defer db.Close()

	// Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Maria Toledo", 11)
	stmt.Exec("José Toledo", 13)
	stmt.Exec("Rafael Toledo", 15)

	// Delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(2000)
	stmt2.Exec(2001)
	stmt2.Exec(2002)
}