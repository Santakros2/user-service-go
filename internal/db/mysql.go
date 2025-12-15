package db

import (
	"database/sql"
	"fmt"
	"log"
	"users-service/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySql(cfg *config.ConfigMySQL) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=false",
		cfg.MySqlUser,
		cfg.MySqlPassword,
		cfg.MySqlHost,
		cfg.MySqlPort,
		cfg.MySqlDB,
	)
	// dsn := "user:pass@tcp(mysql:3306)/users?parseTime=true&tls=false"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		// log.Println("error is ", err)
		// fmt.Scanln()

	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		// log.Println("error is ", err)
		// fmt.Scanln()

	}

	log.Println("connected to MySQL!")
	return db, err
}
