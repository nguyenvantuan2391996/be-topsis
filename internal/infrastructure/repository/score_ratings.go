package repository

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"topsis/internal/domain/model"
)

const (
	BatchSizeCreate = 20
)

type ScoreRatingRepository struct {
	db *gorm.DB
}

func NewScoreRatingRepository(db *gorm.DB) *ScoreRatingRepository {
	return &ScoreRatingRepository{db: db}
}

func (sr *ScoreRatingRepository) BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error {
	return sr.db.CreateInBatches(&scoreRatings, BatchSizeCreate).Error
}

func (sr *ScoreRatingRepository) GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}, sort []string) ([]*model.ScoreRating, error) {
	var result []*model.ScoreRating
	if err := sr.db.Where(queries).Order(strings.Join(sort, " ")).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *ScoreRatingRepository) UpdateScoreRatingWithMap(ctx context.Context, scoreRating *model.ScoreRating, data map[string]interface{}) error {
	if err := sr.db.Model(&scoreRating).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (sr *ScoreRatingRepository) DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error {
	return sr.db.Where(queries).Delete(&model.ScoreRating{}).Error
}
