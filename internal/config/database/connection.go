package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func CreateConnection() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	cfg.DBName = os.Getenv("MYSQL_DATABASE")

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("[DATABASE]Conexão com banco de dados realizada com sucesso", "db", cfg.DBName)

	driver, err := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/config/database/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("[DATABASE] Nenhuma migration pendente", "db", cfg.DBName)
			return
		}
		slog.Error("[DATABASE] Falha ao executar migrations", "erro", err)
		log.Fatal(err)
	}

	slog.Info("[DATABASE] Migration realizada com sucesso", "db", cfg.DBName)
}

func GetConnection() *sql.DB {
	return db
}
