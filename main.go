package main

import (
	"fmt"
	//"os"
)

func main() {
	//currentDir, _ := os.Getwd()
	//fmt.Printf("Current working directory: %s\n", currentDir)

	//personspecific data

	// trying with martins data so greying out this one
	//tinaConfig := CSVConfig{
	//	StartRow: 8, DateCol: 1, DescriptionCol: 2, AmountCol: 3,
	//}

	martinConfig := CSVConfig{
		StartRow: 2, DateCol: 6, DescriptionCol: 9, AmountCol: 10,
	}



	// we read in the orginal account file, code base in csvReader.go
	data, err := ReadCSVFile("./2024_martin.csv", martinConfig)
	// insert the actual filename here!
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(data) == 0 {
        fmt.Println("Warning: CSV file is empty!")
        return
    }

    //fmt.Printf("Number of rows read from CSV: %d\n", len(data))
    
    // Only try to print first row if we have data
    /*if len(data) > 0 {
        fmt.Printf("First row: %v\n", data[0])
    */

	// we clean up the file and strip the header and unneccessary data
	cleanData := cleanTransactions(data, martinConfig)

	//testing the result
	//for _, row := range cleanData {
	//	fmt.Println(row)
	//}

	// here come the data by monthöthing<
	//filterByMonth(cleaned [][]string, months []int) [][]string {
	timeRangeData := filterByMonth(cleanData, []int{1, 3})
	//testing the result
	//fmt.Println("test for correct period")
	//for _, row := range timeRangeData {
	//	fmt.Println(row)
	//}
	//fmt.Println("test for correct period finished")

	// now we start with the calculations (code base in calculator.go)

	//MonthlyCosts := calcTotalExpenses(cleanData)
	//sum of all costs

	costsByCategories := categorizeExpenses(timeRangeData)
	// costs divided by category

	// Get your quality metrics
	//totalTransactions, matchedTransactions, totalSum, matchedSum := calculateReportQuality(cleanData)

	totalTransactions, matchedTransactions, totalCosts, totalIncome, matchedSum := calculateQualityIncomeCosts(timeRangeData)

	months := []int{2, 3}
	//timeRangeData := filterByMonth(cleanData, months)
	printReport(costsByCategories, timeRangeData,
		totalTransactions, matchedTransactions, 
		totalCosts, totalIncome, matchedSum,
		months, 2024)  
		// need to put in year and month, remember!
}
