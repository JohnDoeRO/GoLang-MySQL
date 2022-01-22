package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const username string = "username"
const password string = "password"
const server string = "127.0.0.1"
const port string = "3306"
const db_name string = "database name"

type Make struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	//connect to db
	connect, err := sql.Open("mysql", username+":"+password+"@tcp("+server+":"+port+")/"+db_name)
	if err != nil {
		panic(err.Error())
	}
	defer connect.Close()

	// insert multiple rows
	insert, err := connect.Query("INSERT INTO make (name) VALUES ('Bmw'), ('Citroen'), ('Fiat'), ('Audi')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// single row insert
	// insert, err := connect.Query("INSERT INTO make (name) VALUES ('Bmw')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// update, err := connect.Query("UPDATE make SET name = 'Renault' WHERE id = 1")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer update.Close()

	// delete, err := connect.Query("DELETE FROM make WHERE id = 2")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer delete.Close()

	results, err := connect.Query("SELECT id, name FROM make")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var make Make
		err = results.Scan(&make.Id, &make.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(make.Name)
	}
	defer results.Close()
}
