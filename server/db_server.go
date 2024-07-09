package server

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	driver := config.GetString("DB_DRIVER")
	connectionString := config.GetString("DB_CONN")
	if connectionString == "" {
		log.Fatal("Database connection string is missing")
	}
	dbHandler, err := sql.Open(driver, connectionString)
	if err != nil {
		dbHandler.Close()
		log.Fatalf("Error while validating database: %v", err)
	}
	return dbHandler
}
