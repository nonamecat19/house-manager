package json

import (
	"github.com/google/uuid"
	db "main/db/generated"
	"time"
)

type ItemJson struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt *string   `json:"created_at,omitempty"`
}

func GetItem(item db.Item) ItemJson {
	var createdAt *string
	if item.CreatedAt.Valid {
		formattedTime := item.CreatedAt.Time.Format(time.RFC3339)
		createdAt = &formattedTime
	}

	return ItemJson{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: createdAt,
	}
}

func GetItemList(items []db.Item) []ItemJson {
	var response []ItemJson
	for _, item := range items {
		response = append(response, GetItem(item))
	}
	return response
}
