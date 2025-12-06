package main

import (
	"context"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	inmemRepository := repository.NewInmemRepository()
	svc := service.NewService(inmemRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /preview", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received trip preview request in trip service")
		trip := &domain.RideFareModel{
			ID:                primitive.NewObjectID(),
			UserID:            "user123",
			PackageSlug:       "standard",
			TotalPriceInCents: 1500,
		}

		t, _ := svc.CreateTrip(context.Background(), trip)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t.ID.Hex()))
	})

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}
	log.Println("Starting Trip Service on :8083")
	log.Fatal(server.ListenAndServe())
}
