package my

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DbConnection *sql.DB

func checkDb() {
	DbConnection, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/")
	if err != nil {
		panic(err)
	}
	defer DbConnection.Close()

	_, err = DbConnection.Exec("CREATE DATABASE IF NOT EXISTS my_database")
	if err != nil {
		panic(err)
	}
	DbConnection.Close()

	return
}

// Migrate program.
func Migrate() {
	checkDb()

	// DBに接続
	db, er := gorm.Open("mysql", "root:password@tcp(mysql:3306)/my_database")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})
}
