package model

type Hotel struct {
	ID          int `db:"id"`
	Name        string
	Description string
	Facilities  string
	Address     string
	Price       int
	Rating      float64
}
