package models

type Movie struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}
