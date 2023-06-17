package model

type Review struct {
	ID       int `db:"id"`
	HotelID  int `db:"hotel_id"`
	Score    int
	Feedback string
}
