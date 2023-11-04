package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jcasanella/chat_app/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB

func CreateConnection(cf *config.ConfigValues) {
	// Create Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cf.Host, cf.Port, cf.Username, cf.Password, cf.Database)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	// Reference: https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxIdleConns(20)                 // Connect to be retained before to create a new one
	db.SetMaxOpenConns(100)                // Total Connect = idle + use
	db.SetConnMaxIdleTime(5 * time.Minute) //  Max time a connection can be idle
	db.SetConnMaxLifetime(5 * time.Minute) // Time before close and reopen again a connetion

	orm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	checkError(err)

	db, err = orm.DB()
	checkError(err)

	// Close database
	defer db.Close()

	// check db
	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")
}

func GetDb() *sql.DB {
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
