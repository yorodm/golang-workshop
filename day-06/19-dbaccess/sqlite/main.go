package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	defer db.Close() // Cerrar la base de datos siempre
	// Insertar
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("pparker", "Avengers", "2021-01-01")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// Actualizar
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("spiderman", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// Consultar
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() // Close libera recursos del iterador

	// Eliminar
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
