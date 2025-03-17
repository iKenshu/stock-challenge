package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"stock-challenge/models"
)

func TestFetchStocks(t *testing.T) {
	mockResponse := models.APIResponse{
		Items: []models.Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				Brokerage:  "Brokerage1",
				Action:     "Buy",
				RatingFrom: "Hold",
				RatingTo:   "Buy",
				TargetFrom: "$150",
				TargetTo:   "$200",
				Time:       parseTime("2025-03-01T12:00:00Z")},
			{
				Ticker:     "GOOGL",
				Company:    "Alphabet Inc.",
				Brokerage:  "Brokerage2",
				Action:     "Sell",
				RatingFrom: "Buy",
				RatingTo:   "Sell",
				TargetFrom: "$2500",
				TargetTo:   "$2000",
				Time:       parseTime("2025-03-01T12:00:00Z")},
		},
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseBody)
	}))

	defer server.Close()

	apiURL := server.URL
	apiToken := "test_token"

	stocks, err := FetchStocks(apiURL, apiToken)
	if err != nil {
		t.Fatalf("FetchStocks failed: %v", err)
	}

	if len(stocks) != len(mockResponse.Items) {
		t.Errorf("Expected %d stocks, got %d", len(mockResponse.Items), len(stocks))
	}

	for i, stock := range stocks {
		if stock != mockResponse.Items[i] {
			t.Errorf("Expected stock %d to be %v, got %v", i, mockResponse.Items[i], stock)
		}
	}
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
