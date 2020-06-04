package main

import (
	"./buffers"
	"./operator"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitializeMiddleware()
}

func InitializeMiddleware() {
	/*
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// /* ------------------------- */

	operator.Initialize()
	buffers.Initialize()

	ne := operator.NotificationEngine{ServerHost: "localhost", ServerPort: 3993}
	go ne.Initialize()

	fmt.Scanln()
}
