package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/amy/shetalks/mysql"
	"github.com/amy/shetalks/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// need to abstract DB username and password to config. Its okay for now. Just localhost.
	db, err := sql.Open("mysql", "amy:password@/shetalks")

	if err != nil {
		fmt.Printf("Database opening error: %v", err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Unable to open connection: %v", err)
	}

	// Create Services
	es := &mysql.EventService{DB: db}

	es.Delete(40)

	// Create Router
	router := routes.NewRouter(es)

	// Start API server
	log.Fatal(http.ListenAndServe(":8080", router))

}
