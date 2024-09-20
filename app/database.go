package app

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fanialfi/go-imageVault/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnection *sql.DB
	once         sync.Once
)

func connection() *sql.DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error load config with message : %s\n", err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", *cfg.GetDbUser(), *cfg.GetDbPasswd(), *cfg.GetDbProtocol(), *cfg.GetDbHost(), *cfg.GetDbPort(), *cfg.GetDbName())

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error when create connection to database : %s\n", err.Error())
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(1 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatalf("error when pinging to database : %s\n", err.Error())
	}

	log.Println("Database connected successfully!")
	return db
}

func GetDbInstance() *sql.DB {
	once.Do(func() {
		dbConnection = connection()
	})

	return dbConnection
}

func GetDbInstance1() *sql.DB {
	if dbConnection == nil {
		dbConnection = connection()
	}
	return dbConnection
}
