package repository

import (
	"SoftDesignColloc/internal/database"
	"SoftDesignColloc/internal/model"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type ReviewsRepo struct {
	db database.Queryable
}

func NewReviewsRepo(db database.Queryable) *ReviewsRepo {
	return &ReviewsRepo{db: db}
}

func (r *ReviewsRepo) Create(ctx context.Context, newReview model.Review) (int, error) {
	query := database.PSQL.
		Insert(database.TableReview).
		Columns(
			"hotel_id",
			"score",
			"feedback",
		).
		Values(
			newReview.HotelID,
			newReview.Score,
			newReview.Feedback,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание Review: %w", err)
	}

	return id, nil
}

func (r *ReviewsRepo) ReadReviewByHotelID(ctx context.Context, hotelID int) (*[]model.Review, error) {
	query := database.PSQL.
		Select("id",
			"hotel_id",
			"score",
			"feedback",
		).
		From(database.TableReview).
		Where(squirrel.Eq{
			"hotel_id": hotelID,
		})

	var u []model.Review
	err := r.db.Select(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение Reviews for Hotel: %w", err)
	}
	return &u, nil
}
