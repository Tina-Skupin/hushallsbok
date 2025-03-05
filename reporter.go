package main

import (
	"fmt"
	"strconv"
)

func printReport(totals map[string]float64, transactions [][]string,
	total, matched int, totalCosts, totalIncome, matchedSum float64) {
	PrintReportHeader()
	PrintReportQuality(total, matched, totalCosts, matchedSum)
	PrintFinancialSummary(totalCosts, totalIncome, transactions)
	PrintCategoryBreakdown(totals)
	PrintTransactionDetails(transactions)
}

func PrintReportHeader() {
	fmt.Println("\nBudget Report")
	fmt.Println("=============")
	fmt.Println("")
	fmt.Println("")
}

func PrintReportQuality(total, matched int, totalCosts, matchedSum float64) {
	fmt.Printf("Report Quality check:\n")
	fmt.Println("-----------------")

	fmt.Printf("Total transactions: %d\n", total)
	fmt.Printf("Total costs: %.2f kr\n", totalCosts)
	fmt.Printf("Categorized transactions: %d (%.1f%%)\n",
		matched, float64(matched)/float64(total)*100)
	fmt.Printf("Categorized amount: %.2f kr (%.1f%%)\n",
		matchedSum, matchedSum/totalCosts*100)
	fmt.Println("")
	fmt.Println("")
}

func PrintFinancialSummary(totalCosts float64, totalIncome float64, transactions [][]string) {
	// Basic stats
	fmt.Printf("Summary:\n")
	fmt.Println("-----------------")
	fmt.Printf("Transactions: %d\n", len(transactions))
	fmt.Printf("\nTotal costs: %.2f kr", totalCosts)
	fmt.Printf("\nTotal income: %.2f kr", totalIncome)
	fmt.Printf("\nBalance: %.2f kr", totalIncome-totalCosts)
	fmt.Println("")
	fmt.Println("")
}

func PrintCategoryBreakdown(totals map[string]float64) {
	// Categories
	fmt.Println("\nSpending by Category")
	fmt.Println("-----------------")
	for category, amount := range totals {
		fmt.Printf("%-15s %10.2f kr\n", category+":", amount)
	}
}

// The %-15s creates a left-aligned field 15 chars wide
// %10.2f creates a right-aligned field 10 chars wide with 2 decimal places

func PrintTransactionDetails(transactions [][]string) {
	fmt.Println("\nAnnex. Transaction details")
	fmt.Println("-----------------")
	for _, transaction := range transactions {
		amount, _ := strconv.ParseFloat(transaction[2], 64) // Convert amount back to float
		fmt.Printf("%-12s %-30s %10.2f kr\n",
			transaction[0], // date
			transaction[1], // description
			amount)         // amount
	}
}
