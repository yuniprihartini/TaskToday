package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func ReadEnv(key, defaultValue string) string {
	viper.SetConfigName("configs")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v. not found, default key value : %v, has been loaded", key, defaultValue)
		return defaultValue
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid Type Assertion")
	}
	return value
}

func ConnectionDB() (*sql.DB, error) {

	dbUser := ReadEnv("dbUser", "root")
	dbPass := ReadEnv("dbPass", "Yunip12345#")
	dbHost := ReadEnv("dbHost", "localhost")
	dbPort := ReadEnv("dbPort", "3306")
	dbName := ReadEnv("dbName", "bengkel")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	//fmt.Println("Environment has been activated!")

	return db, nil
}
