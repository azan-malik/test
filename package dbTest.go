package dbTest

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Kcluster struct {
	ID   int    `json:"id"`
	Name string `json:"k_name"`
}

func Dbcon() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected successfully")
	}

	_, err = db.Exec("CREATE DATABASE kdb")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}
}

func Dbselect() {
	_, err = db.Exec("USE kdb")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
}
func Dbtable() {
	stmt, err := db.Prepare("CREATE Table Kcluster(id int NOT NULL AUTO_INCREMENT, k_name varchar(50), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

	defer db.Close()
}
func Dbins() {

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO kcluster VALUES ( 1, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully inserted.")
	}
	defer insert.Close()
}
func Dbsh() {

	// perform a db.Query insert
	// Execute the query
	results, err := db.Query("SELECT id, k_name FROM Kcluster")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var Kcluster Kcluster
		// for each row, scan the result into our tag composite object
		err = results.Scan(&Kcluster.ID, &Kcluster.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(Kcluster.ID)
		log.Printf(Kcluster.Name)
	}
}
