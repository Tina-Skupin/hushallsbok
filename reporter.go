package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
)

func printReport(totals map[string]float64, transactions [][]string,
	total, matched int, totalCosts, totalIncome, matchedSum float64, months []int, year int) {
	// create report filename
	filename := fmt.Sprintf("fianancial_report%d_%02d.csv", year, months[0])

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	PrintReportHeader(file, months, year)
	PrintReportQuality(file, total, matched, totalCosts, matchedSum)
	PrintFinancialSummary(file, totalCosts, totalIncome, transactions)
	PrintCategoryBreakdown(file, totals)
	PrintComment(file)
	PrintTransactionDetails(file, transactions)
}

func PrintReportHeader(w *os.File, months []int, year int) {
	fmt.Fprintf(w, "\nBudget Report for ")
	for i, month := range months {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "%02d", month) // %02d ensures single-digit months get a leading zero
	}
	fmt.Fprintf(w, " %d\n", year)
	fmt.Fprintln(w, "----------------------------------------")
	fmt.Fprintln(w, "=============")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func PrintReportQuality(w *os.File, total, matched int, totalCosts, matchedSum float64) {
	fmt.Fprintf(w, "Report Quality check:\n")
	fmt.Fprintln(w, "-----------------")

	fmt.Fprintf(w, "Total transactions:; %d\n", total)
	fmt.Fprintf(w, "Total costs: %.2f kr\n", totalCosts)
	fmt.Fprintf(w, "Categorized transactions: %d (%.1f%%)\n",
		matched, float64(matched)/float64(total)*100)
	fmt.Fprintf(w, "Categorized amount: %.2f kr (%.1f%%)\n",
		matchedSum, matchedSum/totalCosts*100)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func PrintFinancialSummary(w *os.File, totalCosts float64, totalIncome float64, transactions [][]string) {
	// Basic stats
	fmt.Fprintf(w, "Summary:\n")
	fmt.Fprintln(w, "-----------------")
	fmt.Fprintf(w, "Transactions: %d", len(transactions))
	fmt.Fprintf(w, "\nTotal costs: %.2f kr", totalCosts)
	fmt.Fprintf(w, "\nTotal income: %.2f kr", totalIncome)
	fmt.Fprintf(w, "\nBalance: %.2f kr", totalIncome-totalCosts)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

// TODO: Consider adding category validation if invalid categories become a problem
// Current behavior: Invalid categories are silently ignored in the report
// Potential improvement: Add validateCategory() function using orderedCategories
// as the source of truth for valid categories

func PrintCategoryBreakdown(w *os.File, totals map[string]float64) {
	orderedCategories := []string{
		"rent",
		"flat",
		"Electricity",
		"komm/internet",
		"food",
		"restaurant",
		"clothes",
		"health",
		"hobby",
		"transport",
		"travel",
		"other",
		"Johanna Taschengeld",
		"bank costs",
		"internal, ignore",
		// Add all your categories in the order you want them
	}
	// Categories
	fmt.Fprintln(w, "\nSpending by Category")
	fmt.Fprintln(w, "-----------------")
	for _, category := range orderedCategories {
		if amount, exists := totals[category]; exists {
			fmt.Fprintf(w, "%-15s %10.2f kr\n", category+":", amount)
		}
	}
}

func PrintComment(w *os.File) {
	fmt.Fprintln(w, "\nKommentar: ")
	fmt.Fprintln(w, "-----------------")
	fmt.Fprintln(w, "no coments")
}

// The %-15s creates a left-aligned field 15 chars wide
// %10.2f creates a right-aligned field 10 chars wide with 2 decimal places



func PrintTransactionDetails(w *os.File, transactions [][]string) {
	fmt.Fprintln(w, "\nAnnex. Transaction details")
	fmt.Fprintln(w, "-----------------")
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i][3] < transactions[j][3]  // Compare categories
	})
	for _, transaction := range transactions {
		amount, _ := strconv.ParseFloat(transaction[2], 64) // Convert amount back to float
		fmt.Fprintf(w, "%-12s ; %-30s ; %10.2f ; %-15s \n",
			transaction[0], // date
			transaction[1], // description
			amount, //amount
			transaction[3]) // category
	}
}
