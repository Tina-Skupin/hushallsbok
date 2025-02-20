package main

import (
	"fmt"
)

func main() {
	// we read in the orginal account file, code base in csvReader.go
	data, err := ReadCSVFile("testcsv.csv")
	// insert the actual filename here!
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// we clean up the file and strip the header and unneccessary data
	cleanData := cleanTransactions(data)

    //testing the result
	//for _, row := range cleanData {
	//	fmt.Println(row)
	//}


	// now we start with the calculations (code base in calculator.go)

	//MonthlyCosts := calcTotalExpenses(cleanData)
	//sum of all costs

    costsByCategories := categorizeExpenses(cleanData)
	// costs divided by category

	// Get your quality metrics
    //totalTransactions, matchedTransactions, totalSum, matchedSum := calculateReportQuality(cleanData)

	totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum := calculateQualityIncomeCosts(cleanData)

	printReport(costsByCategories, cleanData,
		totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum)
}
