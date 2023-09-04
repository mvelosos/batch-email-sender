package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	emails := []string{"email1@test.com", "email2@test.com"}

	campaign := NewCampaign(name, content, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	emails := []string{"email1@test.com", "email2@test.com"}

	campaign := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	emails := []string{"email1@test.com", "email2@test.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}
