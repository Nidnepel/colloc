package repository

import (
	"SoftDesignColloc/internal/database"
	"SoftDesignColloc/internal/model"
	"context"
	"fmt"
)

type BookingsRepo struct {
	db database.Queryable
}

func NewBookingsRepo(db database.Queryable) *BookingsRepo {
	return &BookingsRepo{db: db}
}

func (r *BookingsRepo) Create(ctx context.Context, booking model.Booking) (int, error) {
	query := database.PSQL.
		Insert(database.TableBookings).
		Columns(
			"hotel_id",
			"arrival_date",
			"departure_date",
		).
		Values(
			booking.HotelID,
			booking.ArrivalDate,
			booking.DepartureDate,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание Booking: %w", err)
	}

	return id, nil
}
