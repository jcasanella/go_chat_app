package database

import (
	"database/sql"
	"fmt"

	"github.com/jcasanella/chat_app/config"
	_ "github.com/lib/pq"
)

func GetConnection(cf *config.ConfigValues) *sql.DB {
	// Create Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cf.Host, cf.Port, cf.Username, cf.Password, cf.Database)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	// Close database
	defer db.Close()

	// check db
	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
