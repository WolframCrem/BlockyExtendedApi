package database

import (
	"BlockyExtendedApi/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectToDatabase() {
	// check which type is selected
	switch config.LoadedConfig.Database.Type {
	case "mysql":
		db, err := sql.Open("mysql", config.LoadedConfig.Database.Target)
		if err != nil {
			log.Fatalln(err)
		}
		err = db.Ping()

		// handle error
		if err != nil {
			log.Fatalln(err)
		}

		DB = db
		fmt.Println("Connected to the database!")
		break
	case "postgres":
		log.Fatalln("Postgres is not yet supported")
	case "sqlite":
		log.Fatalln("sqlite is not yet supported")
	default:
		log.Fatalln(fmt.Sprintf("%s is not a valid database type", config.LoadedConfig.Database.Type))
	}
}
