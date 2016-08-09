package resource

import (
	"database/sql"
	"golang.org/x/net/context"
	"golangsample/library/database"
)

func GetDB(c context.Context) *sql.DB {
	v, ok := c.Value("database").(*sql.DB)
	if !ok {
		return nil
	}
	return v
}

func UseDB(c context.Context) *sql.DB {
	db := GetDB(c)
	if db == nil {
		db = database.UseEngine()
		c = context.WithValue(c, "database", db)
	}
	return db
}
