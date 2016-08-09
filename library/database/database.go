package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var engine *sql.DB

func UseEngine() *sql.DB {
	if engine == nil {
		engine = initEngine()
	}
	return engine
}

func initEngine() *sql.DB {
	db, err := sql.Open("mysql", "dev:dev@/golangsampledb")
	if err != nil {
		panic(err)
	}
	return db
}
