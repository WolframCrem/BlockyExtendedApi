package database

import (
	"BlockyExtendedApi/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// check which type is selected
	switch config.LoadedConfig.Database.Type {
	case "mysql":
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       config.LoadedConfig.Database.Target,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{})
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
