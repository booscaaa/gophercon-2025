package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Initialize() *sqlx.DB {
	databaseUrl := viper.GetString("database.url")
	db, err := sqlx.Connect("pgx", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(250)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
