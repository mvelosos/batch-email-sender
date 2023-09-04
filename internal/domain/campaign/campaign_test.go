package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body"
	emails  = []string{"email1@test.com", "email2@test.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_ValidatesName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_ValidatesContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_ValidatesEmails(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("emails are required", err.Error())
}
