package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func Initialize() *sqlx.DB {
	databaseUrl := viper.GetString("database.url")
	db, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10)

	return db
}
