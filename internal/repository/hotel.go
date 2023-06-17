package repository

import (
	"SoftDesignColloc/internal/database"
	"SoftDesignColloc/internal/model"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type HotelsRepo struct {
	db database.Queryable
}

func NewHotelsRepo(db database.Queryable) *HotelsRepo {
	return &HotelsRepo{db: db}
}

func (r *HotelsRepo) Read(ctx context.Context, hotelID int) (*model.Hotel, error) {
	query := database.PSQL.
		Select("id",
			"name",
			"description",
			"facilities",
			"address",
			"price",
			"rating",
		).
		From(database.TableHotel).
		Where(squirrel.Eq{
			"id": hotelID,
		})

	var u model.Hotel
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение Hotel: %w", err)
	}
	return &u, nil
}

func (r *HotelsRepo) ReadAll(ctx context.Context) ([]*model.Hotel, error) {
	var items []*model.Hotel
	query := database.PSQL.Select("id",
		"name",
		"address",
		"price",
		"rating",
	).From(database.TableHotel)
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("получение all Hotels: %w", err)
	}
	return items, nil
}
