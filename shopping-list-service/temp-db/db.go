package temp_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	queries "main/queries"
)

func run() {
	connStr := "postgres postgresql://admin:admin@localhost:15432/shopping-list"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	query := queries.New(db)

	items, err := query.ListItems(ctx)

	fmt.Println(items)
}
