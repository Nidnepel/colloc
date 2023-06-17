package repository

import (
	"SoftDesignColloc/internal/database"
	"SoftDesignColloc/internal/model"
	"context"
)

type Hotel interface {
	Read(ctx context.Context, hotelID int) (*model.Hotel, error)
	ReadAll(ctx context.Context) ([]*model.Hotel, error)
}

type Review interface {
	Create(ctx context.Context, newReview model.Review) (int, error)
	ReadReviewByHotelID(ctx context.Context, hotelID int) (*[]model.Review, error)
}

type Booking interface {
	Create(ctx context.Context, booking model.Booking) (int, error)
}

type Repository struct {
	Hotel
	Review
	Booking
}

func NewRepository(db database.Queryable) *Repository {
	return &Repository{
		Hotel:   NewHotelsRepo(db),
		Review:  NewReviewsRepo(db),
		Booking: NewBookingsRepo(db),
	}
}
