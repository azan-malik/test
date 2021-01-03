package dbtesting

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func cdbco() {
	fmt.Print("Hello")
	v := viper.New()
	v.SetConfigFile("database.yaml")
	v.ReadInConfig()
	fmt.Println(v.GetInt("Port"))
	port := v.GetInt("Port")
	fmt.Println(port)
	Username := v.GetString("Username")
	fmt.Println(Username)
	Password := v.GetString("Password")
	fmt.Println(Password)
	Database := v.GetString("Database")
	fmt.Println(Database)
	Host := v.GetString("Host")
	fmt.Println(Host)
	strq := Username + ":" + Password + "@tcp(" + Host + ":" + strconv.Itoa(port) + ")/"

	db, err := sql.Open("mysql", strq)
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
