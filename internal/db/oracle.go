package db

import (
	"database/sql"
	"log"
	"users-service/internal/config"

	// _ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
)

func ConnectOracle(cfg *config.ConfigOracle) (*sql.DB, error) {
	// dsn := fmt.Sprintf(`user="%s" password="%s" connecting="%s:%s%s"`,
	// 	cfg.OracleUser,
	// 	cfg.OraclePassword,
	// 	cfg.OracleHost,
	// 	cfg.OraclePort,
	// 	cfg.OracleService,
	// )

	db, err := sql.Open("oracle", "oracle://user:pass@localhost:1521/XEPDB1")

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to Oracle DB ")
	return db, nil
}
