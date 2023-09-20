package models

type CommonInfoDict struct {
	ID          string `json:"id" bson:"_id"`
	Type        int    `json:"type" bson:"type"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
