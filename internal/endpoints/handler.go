package endpoints

import "batch-email-service/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
