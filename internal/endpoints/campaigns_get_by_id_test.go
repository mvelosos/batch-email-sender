package endpoints

import (
	"batch-email-service/internal/contract"
	internalmock "batch-email-service/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetBy_shouldReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "123",
		Name:    "testing",
		Content: "Hey guys",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetBy_shouldReturnErrorWhenSomethingWrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(res, req)

	assert.Equal(errExpected.Error(), errReturned.Error())
}
