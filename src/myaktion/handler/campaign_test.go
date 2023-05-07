package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nidigeser/go-myaktion/src/myaktion/db"
	"github.com/nidigeser/go-myaktion/src/myaktion/handler"
	"github.com/nidigeser/go-myaktion/src/myaktion/model"
	"github.com/stretchr/testify/suite"
)

type CampaignTestSuit struct {
	suite.Suite
	cleanUpDB func()
}

func (suite *CampaignTestSuit) SetupSuite() {
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite: Start of TestSuite")
	suite.cleanUpDB = db.SetupTestDB(suite.T())
}

// this function executes after all tests executed
func (suite *CampaignTestSuit) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite: End of TestSuite")
	suite.cleanUpDB()
}

// this function executes before each test case
func (suite *CampaignTestSuit) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest: Start of Test")
}

// this function executes after each test case
func (suite *CampaignTestSuit) TearDownTest() {
	fmt.Println("-- From TearDownTest: End of Test")
}

func (suite *CampaignTestSuit) TestCreateCampaign() {
	rr := httptest.NewRecorder()
	jsonData := `{"name":"Covid","organizerName":"Martin","donationMinimum":2,"targetAmount":100,"account":{"name":"Martin","bankName":"DKB","number":"123456"}}`
	req := httptest.NewRequest(http.MethodPost, "/dummy-url", bytes.NewBufferString(jsonData))
	req.Header.Set("Content-Type", "application/json")

	handler := http.HandlerFunc(handler.CreateCampaign)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		suite.T().Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var campaign model.Campaign
	err := json.NewDecoder(rr.Body).Decode(&campaign)
	if err != nil {
		suite.T().Errorf("handler returned unexpected body: %v", err)
		return
	}
	if campaign.ID != 1 {
		suite.T().Errorf("handler returned unexpected ID: got %v want %v", campaign.ID, 1)
	}
}

func TestCampaignTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignTestSuit))
}
