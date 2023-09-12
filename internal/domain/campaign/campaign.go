package campaign

import (
	internalerrors "batch-email-service/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending = "Pending"
	Started = "Started"
	Done    = "Done"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacs := make([]Contact, len(emails))
	for idx, email := range emails {
		contacs[idx].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacs,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}

	return nil, err
}
