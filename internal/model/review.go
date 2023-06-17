package model

import "time"

type Booking struct {
	ID            int       `db:"id"`
	HotelID       int       `db:"hotel_id"`
	ArrivalDate   time.Time `db:"arrival_date"`
	DepartureDate time.Time `db:"departure_date"`
}
