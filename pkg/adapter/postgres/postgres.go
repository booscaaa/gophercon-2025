package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Initialize() *sqlx.DB {
	databaseUrl := viper.GetString("database.url")
	db, err := sqlx.Open("pgx", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(800)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
