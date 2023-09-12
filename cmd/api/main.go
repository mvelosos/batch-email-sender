package main

import (
	"batch-email-service/internal/domain/campaign"
	"batch-email-service/internal/endpoints"
	"batch-email-service/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	campaignService.Repository.Get()

	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))

	http.ListenAndServe(":3000", r)
}
