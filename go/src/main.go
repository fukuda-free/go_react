package main

import (
	"module1/my"

	_ "github.com/go-sql-driver/mysql"
)

// main program.
func main() {
	my.Migrate()
}
