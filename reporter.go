package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"path/filepath"
)

//create the output directory
func ensureOutputDir(dirPath string) error {
    // Check if directory exists
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        // Create directory if it doesn't exist
        return os.MkdirAll(dirPath, 0755)
    }
    return nil
}


func printReport(totals map[string]float64, transactions [][]string,
	total, matched int, totalCosts, totalIncome, matchedSum float64, matchies float64, months []int, year int) {
	// create report filename

	// Ensure output directory exists
	outputDir := "output" // Use the same directory as in saveAsCSV
	if err := ensureOutputDir(outputDir); err != nil {
    	log.Fatalf("Failed to create output directory: %v", err)
}

	// Create text file with directory path
	txtFilename := filepath.Join(outputDir, fmt.Sprintf("financial_report%d_%02d.txt", year, months[0]))
	file, err := os.Create(txtFilename)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()




	PrintReportHeader(file, months, year)
	PrintFinancialSummary(file, totalCosts, totalIncome, transactions)
	PrintCategoryBreakdown(file, totals)
	Printuncategorized(file, totalCosts, matchies)
	PrintComment(file)
	PrintTransactionDetails(file, transactions)
	PrintReportQuality(file, total, matched, totalCosts, matchedSum)

	// Now also save as CSV
	err = saveAsCSV(transactions, totals, year, months, totalCosts, totalIncome, matchedSum)
	//err = saveAsCSV(transactions, totals, year, months)
	if err != nil {
		log.Fatalf("failed to create CSV file: %v", err)
	}

	fmt.Println("Reports generated successfully: ")
	fmt.Println("- Text report:", txtFilename)
	fmt.Println("- CSV report: financial_report" + strconv.Itoa(year) + "_" + fmt.Sprintf("%02d", months[0]) + ".csv")
}

// Add this new function to save as CSV
func saveAsCSV(transactions [][]string, totals map[string]float64, year int, months []int, totalCosts float64, totalIncome float64, matchedSum float64) error {
	//func saveAsCSV(transactions [][]string, totals map[string]float64, year int, months []int) error {
	// Ensure output directory exists
	outputDir := "output" // You can change this to any directory name you prefer
	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}
	// Create CSV filename with directory path
	filename := filepath.Join(outputDir, fmt.Sprintf("financial_report%d_%02d.csv", year, months[0]))		

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add BOM for Excel to recognize UTF-8
	file.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// First write financial summary at the top
	writer.Write([]string{"Financial Summary", "", "", ""})
	writer.Write([]string{"Transactions", fmt.Sprintf("%d", len(transactions)), "", ""})
	writer.Write([]string{"Total income", fmt.Sprintf("%.2f", totalIncome), "", ""})
	writer.Write([]string{"Total costs", fmt.Sprintf("%.2f", totalCosts), "", ""})
	writer.Write([]string{"Balance", fmt.Sprintf("%.2f", totalIncome-totalCosts), "", ""})

	// Write a blank row as separator
	writer.Write([]string{"", "", "", ""})

	// Write category summary
	writer.Write([]string{"Category Summary", "", "", ""})
	writer.Write([]string{"Category", "Amount", "", ""})

	// Get ordered category list (must match the one in PrintCategoryBreakdown)
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
	}

	// Write category totals
	for _, category := range orderedCategories {
		if amount, exists := totals[category]; exists {
			if err := writer.Write([]string{category, fmt.Sprintf("%.2f", amount), "", ""}); err != nil {
				return err
			}
		}
	}

	writer.Write([]string{"uncategorized:", fmt.Sprintf("%.2f", -(totalCosts + matchedSum)), "", ""})

	// Write a blank row as separator
	writer.Write([]string{"", "", "", ""})

	// Write the header row for transactions
	headers := []string{"Date", "Description", "Amount", "Category"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Sort transactions by category for consistency with your text report
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i][3] < transactions[j][3] // Compare categories
	})

	// Write transactions
	for _, transaction := range transactions {
		// Format amount for numerical interpretation
		amount, err := strconv.ParseFloat(transaction[2], 64)
		if err == nil {
			// Create a copy of the transaction slice to avoid modifying the original
			txCopy := make([]string, len(transaction))
			copy(txCopy, transaction)

			// Format with period as decimal separator
			txCopy[2] = fmt.Sprintf("%.2f", amount)

			if err := writer.Write(txCopy); err != nil {
				return err
			}
		} else {
			// If we can't parse the amount, write the original
			if err := writer.Write(transaction); err != nil {
				return err
			}
		}
	}

	return nil
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

func PrintFinancialSummary(w *os.File, totalCosts float64, totalIncome float64, transactions [][]string) {
	// Basic stats
	fmt.Fprintf(w, "Summary:\n")
	fmt.Fprintln(w, "-----------------")
	fmt.Fprintf(w, "Transactions: %d", len(transactions))
	fmt.Fprintf(w, "\nTotal income: %.2f kr", totalIncome)
	fmt.Fprintf(w, "\nTotal costs: %.2f kr", totalCosts)
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
		//"internal, ignore",
		// Add all your categories in the order you want them
	}
	// Categories
	fmt.Fprintln(w, "\nSpending by Category")
	fmt.Fprintln(w, "-----------------")
	for _, category := range orderedCategories {
		if amount, exists := totals[category]; exists {
			fmt.Fprintf(w, "%-15s %17.2f kr\n", category+":", amount)
			//fmt.Fprintf(w, "%-15s %10.2f kr\n", category+":", amount)
		}
	}
}

func Printuncategorized(w *os.File, totalCosts, matchedSum float64) {
	fmt.Fprintf(w, "%-15s %17.2f kr\n", "uncategorized:", -(totalCosts+matchedSum))
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func PrintComment(w *os.File) {
	fmt.Fprintln(w, "\nKommentar: ")
	fmt.Fprintln(w, "-----------------")
	fmt.Fprintln(w, "Einkuenfte uncategorized (Lohn, Barnbidrag)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

// The %-15s creates a left-aligned field 15 chars wide
// %10.2f creates a right-aligned field 10 chars wide with 2 decimal places

func PrintTransactionDetails(w *os.File, transactions [][]string) {
	fmt.Fprintln(w, "\nAnnex. Transaction details")
	fmt.Fprintln(w, "-----------------")
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i][3] < transactions[j][3] // Compare categories
	})
	for _, transaction := range transactions {
		amount, _ := strconv.ParseFloat(transaction[2], 64) // Convert amount back to float
		fmt.Fprintf(w, "%-12s %-30s %10.2f %-15s \n",
			transaction[0], // date
			transaction[1], // description
			amount,         //amount
			transaction[3]) // category
	}
}

func PrintReportQuality(w *os.File, total, matched int, totalCosts, matchedSum float64) {
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Annex. Report Quality check:\n")
	fmt.Fprintln(w, "-----------------")

	fmt.Fprintf(w, "Total transactions: %d\n", total)
	fmt.Fprintf(w, "Total costs: %.2f kr\n", totalCosts)
	fmt.Fprintf(w, "Categorized transactions: %d (%.1f%%)\n",
		matched, float64(matched)/float64(total)*100)
	fmt.Fprintf(w, "Categorized amount: %.2f kr (%.1f%%)\n",
		matchedSum, matchedSum/totalCosts*100)
	fmt.Fprintf(w, "uncategorized: -%.2f kr\n",
		totalCosts+matchedSum)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}
