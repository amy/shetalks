package main

import (
	"database/sql"

	"github.com/amy/shetalks/mysql"
)

func main() {

	// need to abstract DB username and password to config. Its okay for now. Just localhost.
	var db *sql.DB

	mysql.InitDB(db)
}
