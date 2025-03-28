package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CSVConfig struct {
	StartRow       int
	DateCol        int
	AmountCol      int
	DescriptionCol int
}

func ReadCSVFile(filename string, config CSVConfig) ([][]string, error) {
	// open the rootfile
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// read the file
	reader := csv.NewReader(file)

	// Make the reader more flexible
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	reader.LazyQuotes = true    // Be more tolerant of quotes

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}
	return records, nil
}

//cleans up the data, removes whitespaces etc.
func cleanTransactions(records [][]string, config CSVConfig) [][]string {
	var cleaned [][]string

	// ignore the first config lines
	for i := config.StartRow; i < len(records); i++ {
		row := records[i]
		if len(row) == 0 {
			continue
		}

		// Clean up each field, removing extra spaces
		date := strings.TrimSpace(row[config.DateCol])
		description := strings.TrimSpace(row[config.DescriptionCol])
		amount := strings.TrimSpace(row[config.AmountCol])

		// Skip any rows where amount isn't a valid number
		_, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Printf("Skipping row %d due to invalid amount: %s\n", i, amount)
			continue
		}

		cleanedRow := []string{date, description, amount, "Uncategorized"}
		cleaned = append(cleaned, cleanedRow)
	}
	return cleaned
}

// read the first file into the system
func processOneFile(filename string, config CSVConfig) ([][]string, error) {
	records, err := ReadCSVFile(filename, config)
	if err != nil {
		return nil, fmt.Errorf("error processing %s: %v", filename, err)
	}

	cleanedRecords := cleanTransactions(records, config)
	return cleanedRecords, nil
}

// add the second file into the system
func combineTransactions(files []string, configs []CSVConfig) ([][]string, error) {
	var allTransactions [][]string

	for i, file := range files {
		transactions, err := processOneFile(file, configs[i])
		if err != nil {
			return nil, err
		}
		allTransactions = append(allTransactions, transactions...)
	}

	return allTransactions, nil
}

// filters out the relevant time
func filterByMonth(cleaned [][]string, months []int) [][]string {
	// neue datenbank zum füllen
	var transactionBase [][]string

	for _, cleanedRow := range cleaned {

		// If date is "2024-12-22"
		dateParts := strings.Split(cleanedRow[0], "-")
		// dateParts would be ["2024", "12", "22"]
		// The month is dateParts[1]
		monthNum, err := strconv.Atoi(dateParts[1])
		if err != nil {
			continue
		}
		for _, month := range months {
			if monthNum == month {
				transactionBase = append(transactionBase, cleanedRow)
			}
		}
	}
	return transactionBase
}

func filterExclusions(transactions [][]string) [][]string {
	var filtered [][]string

	for _, transaction := range transactions {
		if !shouldExclude(transaction, 1) { // 1 is the description column in cleaned data
			filtered = append(filtered, transaction)
		}
	}
	return filtered
}

// if there are individual transactions you want not to be in the report you can manually exclude them here
func shouldExclude(transaction []string, descriptionCol int) bool {
	description := transaction[descriptionCol]

	exclusionTerms := []string{
		"verf Mobil",
		"via internet",
		"tina",
		"överf",
		"internet",
		"stuff",
		"haushaltsgeld",
		// Add your specific terms here (check by analysis)
	}

	for _, term := range exclusionTerms {
		if strings.Contains(strings.ToLower(description), strings.ToLower(term)) {
			return true
		}
	}
	return false
}
