package main

import (
	"context"
	"fmt"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	inmemRepository := repository.NewInmemRepository()
	svc := service.NewService(inmemRepository)

	trip := &domain.RideFareModel{
		ID:                primitive.NewObjectID(),
		UserID:            "user123",
		PackageSlug:       "standard",
		TotalPriceInCents: 1500,
	}

	t, _ := svc.CreateTrip(context.Background(), trip)

	fmt.Println(t)
}
