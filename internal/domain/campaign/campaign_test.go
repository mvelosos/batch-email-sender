package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body hey!"
	emails  = []string{"email1@test.com", "email2@test.com"}
	fake    = faker.New()
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

func Test_NewCampaign_MustStatusStartAsPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Equal(campaign.Status, Pending)
}

func Test_NewCampaign_ValidatesNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_ValidatesNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(fake.Lorem().Text(30), content, emails)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_ValidatesContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_ValidatesContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), emails)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_ValidatesContactssMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_ValidatesContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is not a valid email", err.Error())
}
