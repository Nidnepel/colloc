package service

import (
	"SoftDesignColloc/internal/model"
	"SoftDesignColloc/internal/repository"
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

type Service struct {
	Hotel
	Review
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Hotel:   NewHotelUsecase(repos.Hotel),
		Review:  NewReviewUsecase(repos.Review),
		Booking: NewBookingUsecase(repos.Booking),
	}
}
