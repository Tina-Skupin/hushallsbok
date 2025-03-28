package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// create the output directory
func ensureOutputDir(dirPath string) error {
	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create directory if it doesn't exist
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

// Reporter handles report generation and saving
type Reporter struct {
	Summary   FinancialSummary
	OutputDir string
}

// NewReporter creates a reporter instance
func NewReporter(summary FinancialSummary, outputDir string) *Reporter {
	return &Reporter{
		Summary:   summary,
		OutputDir: outputDir,
	}
}

// GenerateTextReport formats the financial summary as a text report
func (r *Reporter) GenerateTextReport(transactions [][]string) string {
	s := r.Summary
	var reportBuilder strings.Builder

	// Header
	reportBuilder.WriteString(fmt.Sprintf("\nBudget Report for "))
	for i, month := range s.Months {
		if i > 0 {
			reportBuilder.WriteString(", ")
		}
		reportBuilder.WriteString(fmt.Sprintf("%02d", month))
	}
	reportBuilder.WriteString(fmt.Sprintf(" %d\n", s.Year))
	reportBuilder.WriteString("----------------------------------------\n")
	reportBuilder.WriteString("=============\n\n\n")

	// Financial Summary
	reportBuilder.WriteString("Summary:\n")
	reportBuilder.WriteString("-----------------\n")
	reportBuilder.WriteString(fmt.Sprintf("Transactions: %d", len(transactions)))
	reportBuilder.WriteString(fmt.Sprintf("\nTotal income: %.2f kr", s.TotalIncome))
	reportBuilder.WriteString(fmt.Sprintf("\nTotal costs: %.2f kr", s.TotalCosts))
	reportBuilder.WriteString(fmt.Sprintf("\nBalance: %.2f kr", s.TotalIncome-s.TotalCosts))
	reportBuilder.WriteString("\n\n")

	// Category Breakdown
	reportBuilder.WriteString("\nSpending by Category\n")
	reportBuilder.WriteString("-----------------\n")
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
	}

	for _, category := range orderedCategories {
		if amount, exists := s.TotalsByCategory[category]; exists {
			reportBuilder.WriteString(fmt.Sprintf("%-15s %17.2f kr\n", category+":", amount))
		}
	}

	// Uncategorized
	reportBuilder.WriteString(fmt.Sprintf("%-15s %17.2f kr\n", "uncategorized:", -(s.TotalCosts + s.TotalAllCategorized)))
	reportBuilder.WriteString("\n\n")

	// Comment
	reportBuilder.WriteString("\nKommentar: \n")
	reportBuilder.WriteString("-----------------\n")
	reportBuilder.WriteString("Einkuenfte uncategorized (Lohn, Barnbidrag)\n\n\n\n")

	// Transaction Details
	reportBuilder.WriteString("\nAnnex. Transaction details\n")
	reportBuilder.WriteString("-----------------\n")

	// Sort transactions by category for consistency
	sortedTransactions := make([][]string, len(transactions))
	copy(sortedTransactions, transactions)
	sort.Slice(sortedTransactions, func(i, j int) bool {
		// First sort by category
		if sortedTransactions[i][3] != sortedTransactions[j][3] {
			return sortedTransactions[i][3] < sortedTransactions[j][3]
		}
		// Then by date
		return sortedTransactions[i][0] < sortedTransactions[j][0]
	})

	// Print transactions grouped by category
	currentCategory := ""
	for _, transaction := range sortedTransactions {
		if transaction[3] != currentCategory {
			currentCategory = transaction[3]
			reportBuilder.WriteString(fmt.Sprintf("\n%s\n", currentCategory))
		}

		date := transaction[0]
		description := transaction[1]
		amountStr := transaction[2]

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			continue
		}

		reportBuilder.WriteString(fmt.Sprintf("%s %-40s %10.2f kr\n",
			date, description, amount))
	}

	return reportBuilder.String()
}

// create the txt report
func SaveTextReport(report string, summary *FinancialSummary) error {
    // Ensure output directory exists
    outputDir := "output" // Same directory as your CSV
    if err := ensureOutputDir(outputDir); err != nil {
        return err
    }
    
    // Create text filename with directory path and proper naming convention
    txtFilename := filepath.Join(outputDir, fmt.Sprintf("financial_report%d_%02d.txt", summary.Year, summary.Months[0]))
    
    // Save the file
    return os.WriteFile(txtFilename, []byte(report), 0644)
}

func GenerateCSVReport(summary *FinancialSummary, transactions [][]string, filename string) error {
	// Ensure output directory exists
	outputDir := "output" // You can change this to any directory name you prefer
	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}
	// Create CSV filename with directory path
	csvfilename := filepath.Join(outputDir, fmt.Sprintf("financial_report%d_%02d.csv", summary.Year, summary.Months[0]))

	file, err := os.Create(csvfilename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add BOM for Excel to recognize UTF-8
	file.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write summary information
	writer.Write([]string{"Financial Summary", "", "", ""})
	writer.Write([]string{"Transactions", fmt.Sprintf("%d", len(transactions)), "", ""})
	writer.Write([]string{"Total income", fmt.Sprintf("%.2f", summary.TotalIncome), "", ""})
	writer.Write([]string{"Total costs", fmt.Sprintf("%.2f", summary.TotalCosts), "", ""})
	writer.Write([]string{"Balance", fmt.Sprintf("%.2f", summary.TotalIncome-summary.TotalCosts), "", ""})

	// Write a blank row as separator
	writer.Write([]string{"", "", "", ""})

	
	// Write category breakdown
	writer.Write([]string{"", "", "", ""})
	writer.Write([]string{"CATEGORY BREAKDOWN", "", "", ""})

	// List of categories in desired order
	orderedCategories := []string{
		"rent", "flat", "Electricity", "komm/internet",
		"food", "restaurant", "clothes", "health",
		"hobby", "transport", "travel", "other",
		"Johanna Taschengeld", "bank costs",
	}

	for _, category := range orderedCategories {
		if amount, exists := summary.TotalsByCategory[category]; exists {
			writer.Write([]string{category, "", fmt.Sprintf("%.2f", amount), ""})
		}
	}

	// Write uncategorized amount
	writer.Write([]string{"uncategorized", "", fmt.Sprintf("%.2f", -(summary.TotalCosts + summary.TotalAllCategorized)), ""})

	// Write a blank row as separator
	writer.Write([]string{"", "", "", ""})

	// Write headers
	headers := []string{"Date", "Description", "Amount", "Category"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write transaction data
	for _, transaction := range transactions {
		date := transaction[0]
		description := transaction[1]
		amount := transaction[2]
		category := transaction[3]

		record := []string{date, description, amount, category}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
