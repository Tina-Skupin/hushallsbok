package main

import (
	"strconv"
	"strings"
)

// categories defined as globals, outside of a function
var categories = map[string]string{
	"LIDL":          "food",
	"TESCO":         "food",
	"COOP":          "food",
	"IKEA":          "flat",
	"RUSTA":         "flat",
	"MIO":           "flat",
	"APOTEKET":      "health",
	"APOTEK":        "health",
	"ICA":           "food",
	"STORSTOCKHO":   "transport",
	"MAX":           "food",
	"SYNOPTIK":      "health",
	"DOLLARSTORE":   "hobby",
	"COMVIQ":        "komm/internet",
	"NEW YORKER":    "clothes",
	"GREKISKA":      "food",
	"FLYGBUSSARN":   "travel",
	"PRINCH":        "hobby",
	"BASERBJUDANDE": "other",
	"RESTAURA":		"food",
	"PENDELTAGEN":	"transport",
	"TANZKURS":		"hobby",
	"BAHNHOF":		"komm/internet",
	"CAPIO":		"health",
	"EL":			"flat",
	"STEAM":		"hobby",
	"STORSTOCKHOLM":"transport",
	"RENT":			"rent",
	"TINA":			"internal, ignore",
	"AUDIBLE":		"hobby",
	"ROSSMANN":		"travel",
	"JUMPYARD": 	"hobby",
	"ZAN":			"food",
	"KJELL":		"flat",
	"DB":			"travel",
	"SCOUTS":		"hobby",
	"KURZGESAGT":	"flat",
	"EUROWINGS":	"travel",
	"LUXAIR":		"travel",
	"BOOT":			"hobby",
	"DISNEY PLUS":	"hobby",
	"STUFF":		"internal, ignore",
	"KONSUM":		"travel",
	"MARKANT":		"travel",
	"TANKSTELLE":	"travel",
	"BANKKORT":		"bank costs",
	"INTENRETBET":	"bank costs",
	"VY BUSS":		"travel",
	"SKANSEN":		"hobby",
	"VAXHOLMS":		"travel",
	"FILMSTADEN":	"hobby",
	"STORSK":		"travel",
	"SAS":			"travel",
	"BLIZZARD":		"hobby",
	"VIA INTERNET":		"internal, ignore",
	"ARLANDA":		"travel",
	"TASCHENGEL":	"Johanna Taschengeld",
	"EVENTIM":		"hobby",
}

func categorizeExpenses(transactions [][]string) map[string]float64 {

	// Initialize totals map
	totals := map[string]float64{
		"rent":          0.0,
		"flat":          0.0,
		"komm/internet": 0.0,
		"food":          0.0,
		"clothes":       0.0,
		"health":        0.0,
		"transport":     0.0,
		"Johanna Taschengeld": 0.0,
		"hobby":         0.0,
		"travel":        0.0,
		"other":         0.0,
		"internal, ignore": 0.0,
		"bank costs":	0.0,
	}

	for _, transaction := range transactions {
		description := strings.ToUpper(transaction[1])
		amount, err := strconv.ParseFloat(transaction[2], 64)
		if err != nil {
			continue // skip this transaction if amount can't be converted
		}

		// Now check if description contains any store name
		for store, category := range categories {
			if strings.Contains(description, store) {
				totals[category] += amount
				break // once we find a match, we can stop checking other stores
			}
		}
	}
	return totals
}

func calculateQualityIncomeCosts(transactions [][]string) (int, int, float64, float64, float64) {
	totalTransactions := len(transactions)
	matchedTransactions := 0
	totalCosts := 0.0
	totalIncome := 0.0
	matchedSum := 0.0

	for _, transaction := range transactions {
		amount, _ := strconv.ParseFloat(transaction[2], 64)
		if amount < 0 {
			totalCosts += -amount // Make positive for display
		} else {
			totalIncome += amount
		}
		description := strings.ToUpper(transaction[1])
		for store := range categories {
			if strings.Contains(description, store) {
				matchedTransactions++
				matchedSum += amount
				break
			}
		}
	}
	return totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum
}
