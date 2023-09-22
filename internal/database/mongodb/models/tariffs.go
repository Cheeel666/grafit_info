package models

import "time"

// Tariff - тарифы и вся инфа по ним.
type Tariff struct {
	CreatedAt   int64     `json:"created_at" bson:"created_at"`
	UpdatedAt   int64     `json:"updated_at" bson:"updated_at"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	DateStart   time.Time `json:"date_start" bson:"date_start"`
	DateEnd     time.Time `json:"date_end" bson:"date_end"`
	Price       float64   `json:"price" bson:"price"`
	Duration    string    `json:"duration" bson:"duration"`
	Type        string    `json:"type" bson:"type"`
}

type Tariffs []*Tariff

type FindReq struct {
	ID   string `json:"id" bson:"_id"`
	Type int    `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}
