package service

import (
	"SoftDesignColloc/internal/model"
	"SoftDesignColloc/internal/repository"
	"context"
)

type BookingUsecase struct {
	repo repository.Booking
}

func NewBookingUsecase(booking repository.Booking) *BookingUsecase {
	return &BookingUsecase{repo: booking}
}

func (u *BookingUsecase) Create(ctx context.Context, newBooking model.Booking) (int, error) {
	return u.repo.Create(ctx, newBooking)
}
