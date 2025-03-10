package main

import (
	"fmt"
	"strconv"
)

func printReport(totals map[string]float64, transactions [][]string,
	total, matched int, totalCosts, totalIncome, matchedSum float64, months []int, year int) {
	PrintReportHeader(months, year)
	PrintReportQuality(total, matched, totalCosts, matchedSum)
	PrintFinancialSummary(totalCosts, totalIncome, transactions)
	PrintCategoryBreakdown(totals)
	PrintTransactionDetails(transactions)
}

func PrintReportHeader(months []int, year int) {
	fmt.Printf("\nBudget Report for ")
	for i, month := range months {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%02d", month) // %02d ensures single-digit months get a leading zero
	}
	fmt.Printf(" %d\n", year)
	fmt.Println("----------------------------------------")
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


// TODO: Consider adding category validation if invalid categories become a problem
// Current behavior: Invalid categories are silently ignored in the report
// Potential improvement: Add validateCategory() function using orderedCategories
// as the source of truth for valid categories

func PrintCategoryBreakdown(totals map[string]float64) {
	orderedCategories := []string{
		"rent",
		"flat",
		"food",
		"komm/internet",
		"clothes",
		"health",
		"transport",
		"Johanna Taschengeld",
		"hobby",
		"travel",
		"other",
		"internal, ignore",
		"bank costs",
		// Add all your categories in the order you want them
	}
	// Categories
	fmt.Println("\nSpending by Category")
	fmt.Println("-----------------")
	for _, category := range orderedCategories {
		if amount, exists := totals[category]; exists {
			fmt.Printf("%-15s %10.2f kr\n", category+":", amount)
		}
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
