package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	return s.repo.CreateTrip(ctx, &domain.TripModel{
		UserID:   fare.UserID,
		Status:   "created",
		RideFare: fare,
	})
}
