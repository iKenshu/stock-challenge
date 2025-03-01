package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStockSerialization(t *testing.T) {
	stock := Stock{
		Ticker:     "AAPL",
		Company:    "Apple Inc.",
		Brokerage:  "Brokerage1",
		Action:     "Buy",
		RatingFrom: "Hold",
		RatingTo:   "Buy",
		TargetFrom: "$150",
		TargetTo:   "$200",
		Time:       parseTime("2025-03-01T12:00:00Z"),
	}

	stockJSON, err := json.Marshal(stock)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var deserializedStock Stock
	err = json.Unmarshal(stockJSON, &deserializedStock)
	if err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if stock != deserializedStock {
		t.Errorf("Stocks are not equal")
	}
}

func TestAPIResponseSerialization(t *testing.T) {
	apiResponse := APIResponse{
		Items: []Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				Brokerage:  "Brokerage1",
				Action:     "Buy",
				RatingFrom: "Hold",
				RatingTo:   "Buy",
				TargetFrom: "$150",
				TargetTo:   "$200",
				Time:       parseTime("2025-03-01T12:00:00Z"),
			},
			{
				Ticker:     "GOOGL",
				Company:    "Alphabet Inc.",
				Brokerage:  "Brokerage2",
				Action:     "Sell",
				RatingFrom: "Buy",
				RatingTo:   "Sell",
				TargetFrom: "$2500",
				TargetTo:   "$2000",
				Time:       parseTime("2025-03-01T12:00:00Z"),
			},
		},
		NextPage: "https://api.example.com/stocks?page=2",
	}

	apiResponseJSON, err := json.Marshal(apiResponse)

	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var deserializedAPIResponse APIResponse
	err = json.Unmarshal(apiResponseJSON, &deserializedAPIResponse)
	if err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if apiResponse.NextPage != deserializedAPIResponse.NextPage {
		t.Errorf("NextPage is not equal")
	}

	if len(apiResponse.Items) != len(deserializedAPIResponse.Items) {
		t.Errorf("Items length is not equal")
	}

	for i := range apiResponse.Items {
		if apiResponse.Items[i] != deserializedAPIResponse.Items[i] {
			t.Errorf("Item %d is not equal", i)
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
