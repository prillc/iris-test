package models

type Movie struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

type MovieHistory struct {
	ID      int64 `json:"id" gorm:"primaryKey"`
	Created int32 `json:"created"`
	MovieID int64 `json:"movie_id"`
	Movie   Movie `json:"movie" gorm:"foreignKey:MovieID"`
}
