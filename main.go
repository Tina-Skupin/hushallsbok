package main

import (
	"fmt"
	"log"
	//"path/filepath"
)

func main() {

	// all the input infos are being put in here:


	// configuration of the raw csv (where in the banc data are the infos we want filtered out)
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

	months := []int{6}
	// time scope
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

	//die alten calls
	//costsByCategories, matchies := categorizeExpenses(finalTransactions)
	// costs divided by category
	//totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum := calculateQualityIncomeCosts(finalTransactions)

	//Report

	//alter Report
	/*printReport(costsByCategories, finalTransactions,
			totalTransactions, matchedTransactions,
			totalCosts, totalIncome, matchedSum, matchies,
			months, 2024)
		fmt.Println("Bericht wurde erstellt")
		fmt.Println("=============")
		// need to put in year and month, remember!
	} */

	//neuer Report

	// Generate and save the text report

	// Create a reporter instance
	reporter := NewReporter(summary, "output") // or whatever output directory you want

	// Generate the text report
	textReport := reporter.GenerateTextReport(finalTransactions)
	fmt.Println("Report length:", len(textReport))
	fmt.Println("First 100 chars:", textReport[:100])

	// Save the text report
	// Assuming outputFolder is where your CSV is saved
	//textFilePath := filepath.Join("output", "financial_report.txt")
	err = SaveTextReport(textReport, &summary)
	//err = SaveTextReport(textReport, textFilePath)
	fmt.Println("SaveTextReport result:", err)
	if err != nil {
		log.Fatalf("Failed to save text report: %v", err)
	}

	// Generate and save the CSV report
	err = GenerateCSVReport(&summary, finalTransactions, "financial_report.csv")
	if err != nil {
		log.Fatalf("Failed to save CSV report: %v", err)
	}
}
