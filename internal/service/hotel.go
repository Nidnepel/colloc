package service

import (
	"SoftDesignColloc/internal/model"
	"SoftDesignColloc/internal/repository"
	"context"
)

type HotelUsecase struct {
	repo repository.Hotel
}

func NewHotelUsecase(hotel repository.Hotel) *HotelUsecase {
	return &HotelUsecase{repo: hotel}
}

func (u *HotelUsecase) Read(ctx context.Context, hotelID int) (*model.Hotel, error) {
	return u.Read(ctx, hotelID)
}

func (u *HotelUsecase) ReadAll(ctx context.Context) ([]*model.Hotel, error) {
	return u.ReadAll(ctx)
}
