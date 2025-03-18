package main

import (
	"strconv"
	"strings"
)

// categories defined as globals, outside of a function
var categories = map[string]string{
	"RENT":           "rent",
	"RENT DELAY":     "rent",
	"IKEA":           "flat",
	"RUSTA":          "flat",
	"MIO":            "flat",
	"KJELL":          "flat",
	"KJELL & CO 14": 	"flat",
	"KURZGESAGT":     "flat",
	"JULA":           "flat",
	"NORMAL":         "flat",
	"CLAS OHLSON":    "flat",
	"BLOMSTERLANDET": "flat",
	"IF SKADE":	"flat",
	"BAHNHOF":        "komm/internet",
	"COMVIQ":         "komm/internet",
	"LIDL":           "food",
	"TESCO":          "food",
	"PRESSBYRAN":     "food",
	"COOP":           "food",
	"SYSTEMBOLAGET":  "food",
	"CHOCO MANIA":    "restaurant",
	"BRODERNAS":      "restaurant",
	"SUBWAY":         "restaurant",
	"ICA":            "food",
	"HEMKOP":         "food",
	"GROSSEN":        "food",
	"MAX":            "restaurant",
	"PLATTAN":        "restaurant",
	"GREKISKA":       "restaurant",
	"KOKUKUJIRA":     "restaurant",
	"RESTAURA":       "restaurant",
	"BING WANG":      "restaurant",
	"ZAN":            "restaurant",
	"MCDONALDS":      "restaurant",
	"ESPRESSO":       "restaurant",
	"MARRYBROWN":     "restaurant",
	"K#LLARBYN":      "restaurant",
	"CHOKLADFABRIKE": "restaurant",
	"HONEYCOMB":      "restaurant",
	"YUME SHUSHI":    "restaurant",
	"HERMANS":        "restaurant",
	"WAYNES COFFEE":	"restaurant",
	"NEW YORKER":     "clothes",
	"ARTIKEL 2":       "clothes",
	"DECATHLON":      "clothes",
	"KAPPAHL":        "clothes",
	"SKYDDS":         "clothes",
	"SKOPUNKTEN":     "clothes",
	"HM":             "clothes",
	"STADIUM":			"clothes",
	"APOTEKET":       "health",
	"APOTEK":         "health",
	"CAPIO":          "health",
	"SYNOPTIK":       "health",
	"KVINNA STOCKHO": "health",
	"EMALJKLINIKEN":  "health",
	"VACCINDIREKT":   "health",
	"BRILLE":         "health",
	"SYNSAM":			"health",
	"PRINCH":         "hobby",
	"TANZKURS":       "hobby",
	"STEAM":          "hobby",
	"PANDURO":        "hobby",
	"AUDIBLE":        "hobby",
	"SCOUTS":         "hobby",
	"UPPT": 	"hobby",
	"BOOT":           "hobby",
	"DISNEY PLUS":    "hobby",
	"FILMSTADEN":     "hobby",
	"DOLLARSTORE":    "hobby",
	"EVENTIM":        "hobby",
	"KLARNA":         "hobby",
	"JUMPYARD":       "hobby",
	"SKANSEN":        "hobby",
	"STIFTELSEN SKAN": "hobby",
	"BLIZZARD":       "hobby",
	"BIBLIOTE":       "hobby",
	"KL#TTERCENTRET": "hobby",
	"J#RVABADET":     "hobby",
	"SF KONGRESSER":	"hobby",
	"FRYSHUSET":		"hobby",
	"AKADEMIBOKHAND":	"hobby",
	"ADLIBRIS":			"hobby",
	"STORSTOCKHOLM":  "transport",
	"PENDELTAGEN":    "transport",
	"STORSTOCKHO":    "transport",
	"SWISH SJ AB":    "transport",
	"SJ APP":         "transport",
	"ROSSMANN":       "travel",
	"NORRA JÄRVA SC": "travel",
	"DB":             "travel",
	"FLYGBUSSARN":    "travel",
	"EUROWINGS":      "travel",
	"LUXAIR":         "travel",
	"KONSUM":         "travel",
	"MARKANT":        "travel",
	"TANKSTELLE":     "travel",
	"ARLANDA":        "travel",
	"VY BUSS":        "travel",
	"VAXHOLMS":       "travel",
	"STORSK":         "travel",
	"SAS":            "travel",
	"BANKKORT":       "bank costs",
	"INTENRETBET":    "bank costs",
	"BASERBJUDANDE":  "bank costs",
	"VIA INTERNET":   "internal, ignore",
	"TINA":           "internal, ignore",
	"STUFF":          "internal, ignore",
	"TASCHENGEL":     "Johanna Taschengeld",
}

func categorizeExpenses(transactions [][]string) map[string]float64 {

	// Initialize totals map
	totals := map[string]float64{
		"rent":                0.0,
		"flat":                0.0,
		"Electricity":         0.0,
		"komm/internet":       0.0,
		"food":                0.0,
		"clothes":             0.0,
		"health":              0.0,
		"transport":           0.0,
		"Johanna Taschengeld": 0.0,
		"hobby":               0.0,
		"travel":              0.0,
		"other":               0.0,
		"internal, ignore":    0.0,
		"bank costs":          0.0,
	}

	for _, transaction := range transactions {
		description := strings.ToUpper(transaction[1])
		amount, err := strconv.ParseFloat(transaction[2], 64)
		if err != nil {
			continue // skip this transaction if amount can't be converted
		}
			// First check for our special "el" cases
		if description == "EL" || strings.Contains(description, "EL NET") {
    		totals["Electricity"] += amount
    		transaction[3] = "Electricity"  // Set category to Electricity
		} else {
    		// Check all other categories
    		for store, category := range categories {
        		if store != "el" && strings.Contains(description, store) {
            		totals[category] += amount
            		transaction[3] = category // Set the category
            		break  // Stop checking once we find a match
				}
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
