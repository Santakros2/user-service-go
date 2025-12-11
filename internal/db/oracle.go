package db

import (
	"database/sql"
	"fmt"
	"log"
	"users-service/internal/config"

	_ "github.com/godror/godror"
)

func ConnectOracle(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(`user="%s" password="%s" connecting="%s:%s%s"`,
		cfg.OracleUser,
		cfg.OraclePassword,
		cfg.OracleHost,
		cfg.OraclePort,
		cfg.OracleService,
	)

	db, err := sql.Open("godror", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to Oracle DB ")
	return db, nil
}
