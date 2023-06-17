package service

import (
	"SoftDesignColloc/internal/model"
	"SoftDesignColloc/internal/repository"
	"context"
)

type ReviewUsecase struct {
	repo repository.Review
}

func NewReviewUsecase(review repository.Review) *ReviewUsecase {
	return &ReviewUsecase{repo: review}
}

func (u *ReviewUsecase) Create(ctx context.Context, newReview model.Review) (int, error) {
	return u.repo.Create(ctx, newReview)
}

func (u *ReviewUsecase) ReadReviewByHotelID(ctx context.Context, hotelID int) (*[]model.Review, error) {
	return u.repo.ReadReviewByHotelID(ctx, hotelID)
}
