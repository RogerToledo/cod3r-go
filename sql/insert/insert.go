package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, errSt := db.Prepare("insert into usuarios(nome) values(?)")
	if errSt != nil {
		panic(errSt)
	}
	//stmt.Exec("Maria")
	//stmt.Exec("João")

	nomes := []string{"Marcelo", "Marcio", "Ieda"}

	var res sql.Result
	var count []int64
	var id int64

	for _, nome := range nomes {
		res, _ = stmt.Exec(nome)
		id, _ = res.LastInsertId()
		fmt.Println(id)
		count = append(count, id)
	}

	fmt.Println(len(count))

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
