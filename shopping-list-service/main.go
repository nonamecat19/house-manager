package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"main/db"
)

func main() {
	fmt.Println("Hello World")
	queries := db.GetQueries()
	ctx := context.Background()
	items, err := queries.ListItems(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("ListItems:", items)
}
