package main

import (
	"fmt"
	"log"
)

func main() {

	// all the input infos are being put here in here:


	// configuration of the raw csv (where in the banc data are the infos we want filtered out, first column = 0)
	configs := []CSVConfig{
		{
			StartRow:       9,
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

	// time scope
	months := []int{7}
	//months := []int{1,2,3,4,5,6,7,8,9,10,11,12} //if several months
	year := 2023

	// Get combined transactions instead of single file
	transactions, err := combineTransactions(files, configs)
	if err != nil {
		log.Fatalf("Error combining transactions: %v", err)
	}

	// Apply filters to get final transaction set
	cleaned := filterByMonth(transactions, months)
	finalTransactions := filterExclusions(cleaned)

	// Analysis
	summary := calculateFinances(finalTransactions, year, months)

	// Create a reporter instance
	reporter := NewReporter(summary, "output") // or whatever output directory you want

	// Generate the text report
	textReport := reporter.GenerateTextReport(finalTransactions)

	// save the text report
	err = SaveTextReport(textReport, &summary)
	fmt.Println("txt Bericht wurde erstellt")
	if err != nil {
		log.Fatalf("Failed to save text report: %v", err)
	}

	// Generate and save the CSV report
	err = GenerateCSVReport(&summary, finalTransactions, "financial_report.csv")
	fmt.Println("csv Bericht wurde erstellt")
	if err != nil {
		log.Fatalf("Failed to save CSV report: %v", err)
	}
}
