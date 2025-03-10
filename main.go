package main

import (
	"log"
)

func main() {

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


    files := []string{
        "2024_tina.csv",   // source file dataset tina 
        "2024_martin.csv",    // source file dataset martin 
    }

	months := []int{11, 12}
    

// Get combined transactions instead of single file

	transactions, err := combineTransactions(files, configs)
	if err != nil {
		log.Fatalf("Error combining transactions: %v", err)
	}
		
	cleaned := filterByMonth(transactions, months)
	finalTransactions := filterExclusions(cleaned)
		
		
	// Analysis

	costsByCategories := categorizeExpenses(finalTransactions)
	// costs divided by category
	totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum := calculateQualityIncomeCosts(finalTransactions)

	//Report
	printReport(costsByCategories, finalTransactions,
		totalTransactions, matchedTransactions, 
		totalCosts, totalIncome, matchedSum,
		months, 2024)  
		// need to put in year and month, remember!
}
