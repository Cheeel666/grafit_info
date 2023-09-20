package models

// DocumentDict - словарь документов и прочего. Например, счет, абонимент и т.д. Отдельная таблица.
type DocumentDict struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Type int    `json:"type" bson:"type"`
}
