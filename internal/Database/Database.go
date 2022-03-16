package Database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	// Create Table
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	// stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urlinfo (`ShortID` string NOT NULL, `Link` string not NULL, `Expireat` int64 not NULL)")
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urlinfo (ShortID string, Link string, Expireat string);")
	ErrCheck(err)
	stmt.Exec()
	// fmt.Print("Table Created")
	ErrCheck(err)
	db.Close()
}

func AddData(ShortID string, Link string, ExpireAt int64) {
	CreateTable()
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	stmt, err := db.Prepare("INSERT INTO urlinfo(ShortID, Link, Expireat) values(?,?,?)")
	ErrCheck(err)
	res, err := stmt.Exec(ShortID, Link, ExpireAt)
	ErrCheck(err)
	res.LastInsertId()
	db.Close()
}

func QueryData(ID string) (bool, string) {
	db, err := sql.Open("sqlite3", "./data.db")
	ErrCheck(err)
	now := time.Now().Unix()
	rows, err := db.Query("SELECT * FROM urlinfo WHERE ShortID = ?", ID)
	ErrCheck(err)
	status := false
	URL := ""
	for rows.Next() {
		var Link string
		var Expireat int64
		var ShortLink string
		err = rows.Scan(&ShortLink, &Link, &Expireat)
		ErrCheck(err)
		if Link != "" {
			URL = Link
			if Expireat > now {
				status = true
			} else {
				status = false
			}
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
