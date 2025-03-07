package db

import (
	"database/sql"
	"log"
	sqlcgen "main/db/generated"
	"sync"

	_ "github.com/lib/pq"
)

var (
	instance *sqlcgen.Queries
	once     sync.Once
	db       *sql.DB
)

func GetQueries() *sqlcgen.Queries {
	once.Do(func() {
		var err error
		connStr := "postgresql://admin:admin@localhost:15432/shopping-list?sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		instance = sqlcgen.New(db)
	})
	return instance
}
