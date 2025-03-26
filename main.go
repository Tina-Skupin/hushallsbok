package main

import (
	"fmt"
	"log"
)

func main() {

	// all the input infos are being put in here:
	// configuration of the raw csv (where in the table are the infos we want filtered out)
	configs := []CSVConfig{
		{
			StartRow:       8,
			DateCol:        1,
			DescriptionCol: 2,
			AmountCol:      3,
		},
		{
			StartRow:       2,
			DateCol:        6,
			DescriptionCol: 9,
			AmountCol:      10,
		},
	}

	//source files
	files := []string{
		"2024_tina.csv",   // source file dataset 1
		"2024_martin.csv", // source file dataset 2
	}

	months := []int{3}
	// time scope
	//months := []int{1,2,3,4,5,6,7,8,9,10,11,12} //if several months

	// Get combined transactions instead of single file
	transactions, err := combineTransactions(files, configs)
	if err != nil {
		log.Fatalf("Error combining transactions: %v", err)
	}

	cleaned := filterByMonth(transactions, months)
	finalTransactions := filterExclusions(cleaned)

	// Analysis

	costsByCategories, matchies := categorizeExpenses(finalTransactions)
	// costs divided by category
	totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum := calculateQualityIncomeCosts(finalTransactions)

	//Report
	printReport(costsByCategories, finalTransactions,
		totalTransactions, matchedTransactions,
		totalCosts, totalIncome, matchedSum, matchies,
		months, 2024)
	fmt.Println("Bericht wurde erstellt")
	fmt.Println("=============")
	// need to put in year and month, remember!
}
