package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	// Create Table
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	// stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urlinfo (`ShortID` string NOT NULL, `Link` string not NULL, `Expireat` string not NULL)")
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urlinfo (ShortID string, Link string, Expireat string);")
	ErrCheck(err)
	stmt.Exec()
	// fmt.Print("Table Created")
	ErrCheck(err)
	db.Close()
}

func AddData(ShortID, Link, ExpireAt string) {
	CreateTable()
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	stmt, err := db.Prepare("INSERT INTO urlinfo(ShortID, Link, Expireat) values(?,?,?)")
	ErrCheck(err)
	res, err := stmt.Exec(ShortID, Link, ExpireAt)
	ErrCheck(err)
	id, err := res.LastInsertId()
	ErrCheck(err)
	fmt.Println(id)
	db.Close()
}

func QueryData(ID string) (bool, string) {
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	rows, err := db.Query("SELECT link FROM urlinfo WHERE ShortID = ?", ID)
	ErrCheck(err)
	status := false
	var URL string
	for rows.Next() {
		var link string
		err = rows.Scan(&link)
		ErrCheck(err)
		if link != "" {
			URL = link
			status = true
		} else {
			URL = ""
			status = false
		}
	}
	return status, URL
}

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}
