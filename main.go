package main

import (
	//"fmt"
	"log"
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
        "2024_tina.csv",   // Use your actual file names
        "2024_martin.csv",    // Use your actual file names
    }
    

	//martinConfig := CSVConfig{
		//StartRow: 2, DateCol: 6, DescriptionCol: 9, AmountCol: 10,
	//}

// Get combined transactions instead of single file
	transactions, err := combineTransactions(files, configs)
	if err != nil {
		log.Fatalf("Error combining transactions: %v", err)
	}

	//cleanData := cleanTransactions(transactions, martinConfig)

	//testing the result
	//for _, row := range cleanData {
	//	fmt.Println(row)
	//}

	// here come the data by monthöthing<
	//filterByMonth(cleaned [][]string, months []int) [][]string {
	timeRangeData := filterByMonth(transactions, []int{11, 12})
	// timeRangeData := filterByMonth(cleanData, []int{1, 4})
	//if i want to filter for several months

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

	months := []int{1, 2}
	//timeRangeData := filterByMonth(cleanData, months)
	printReport(costsByCategories, timeRangeData,
		totalTransactions, matchedTransactions, 
		totalCosts, totalIncome, matchedSum,
		months, 2024)  
		// need to put in year and month, remember!
}
