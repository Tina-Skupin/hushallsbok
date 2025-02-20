package main

import (
	"strconv"
	"strings"
)

// calculates the total expenses in a month
/*func calcTotalExpenses(transactions [][]string) float64 {
	var totalExpenses float64
	for _, transaction := range transactions {
		amount, err := strconv.ParseFloat(transaction[2], 64)
		if err != nil {
			continue
		}
		totalExpenses += amount
	}
	return totalExpenses */


//categories defined as globals, outside of a function
var categories = map[string]string {
	"LIDL":        "food",
		"TESCO":       "food",
		"COOP":        "food",
		"IKEA":        "flat",
		"RUSTA":       "flat",
		"MIO":         "flat",
		"APOTEKET":    "health",
		"ICA":         "food",
		"STORSTOCKHO": "transport",
		"MAX":         "food",
		"SYNOPTIK":    "health",
		"DOLLARSTORE": "hobby",
		"COMVIQ":      "komm/internet",
}

func categorizeExpenses(transactions [][]string) map[string]float64{
	
	// Initialize totals map
	totals := map[string]float64{
		"food":                   0.0,
		"flat":                   0.0,
		"health":                 0.0,
		"transport":              0.0,
		"rent":                   0.0,
		"hobby":                  0.0,
		"komm/internet": 0.0,
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

func calculateQualityIncomeCosts(transactions [][]string) (int, int, float64, float64, float64)  {
    totalTransactions := len(transactions)
    matchedTransactions := 0
    totalCosts := 0.0
	totalIncome := 0.0
    matchedSum := 0.0
    
	for _, transaction := range transactions {
        amount, _ := strconv.ParseFloat(transaction[2], 64)
        if amount < 0 {
            totalCosts += -amount  // Make positive for display
        } else {
            totalIncome += amount
		}
        description := transaction[1]
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