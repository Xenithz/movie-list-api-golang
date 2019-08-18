package mydb

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	Host     string
	Password string
	User     string
	Port     int
	DBName   string
}

func ConnectToDB() *sql.DB {
	godotenv.Load()
	config := dbConfig{}
	setConfig(&config)

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func setConfig(config *dbConfig) {
	config.Host = os.Getenv("DB_HOST")
	if len(config.Host) == 0 {
		panic("host problem")
	}

	config.Password = os.Getenv("DB_PASSWORD")
	if len(config.Password) == 0 {
		panic("password problem")
	}

	config.User = os.Getenv("DB_USER")
	if len(config.User) == 0 {
		panic("user problem")
	}

	portHolder, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if portHolder == 0 {
		panic("port problem")
	}
	if err != nil {
		panic("failed string")
	}
	config.Port = portHolder

	config.DBName = os.Getenv("DB_NAME")
	if len(config.DBName) == 0 {
		panic("dbname problem")
	}
}
