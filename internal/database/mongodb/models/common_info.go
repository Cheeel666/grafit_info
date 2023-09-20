package models

import "time"

// CommonInfo - общая информация на сайте. Все футеры, хедеры и т.д. Режим работы, телефоны и т.д.
type CommonInfo struct {
	ID        string    `json:"id" bson:"_id"`
	Type      int       `json:"type" bson:"type"`
	Name      string    `json:"name" bson:"name"`
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
