package db

import (
	"strconv"
	"strings"
)

func parsePrice(priceStr string) (float64, error) {
	priceStr = strings.TrimPrefix(priceStr, "$")
	return strconv.ParseFloat(priceStr, 64)
}
