package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type CSVConfig struct {
	StartRow       int
	DateCol        int
	AmountCol      int
	DescriptionCol int
}


func ReadCSVFile(filename string, config CSVConfig) ([][]string, error) {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// read the file
	reader := csv.NewReader(file)
    
    // Make the reader more flexible
    reader.FieldsPerRecord = -1  // Allow variable number of fields
    reader.LazyQuotes = true     // Be more tolerant of quotes

    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("error reading CSV: %v", err)
    }

    /*Debug: Print the first few rows
    for i := 0; i < min(3, len(records)); i++ {
        fmt.Printf("Row %d has %d fields: %v\n", i, len(records[i]), records[i])
    } */

    return records, nil
}


func cleanTransactions(records [][]string, config CSVConfig) [][]string {
	var cleaned [][]string

    /*debug fmt.Printf("Starting to clean transactions from row %d\n", config.StartRow)
    fmt.Printf("Using columns: Date=%d, Description=%d, Amount=%d\n", 
        config.DateCol, config.DescriptionCol, config.AmountCol) */

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

		// Print first few rows of cleaned data
        /*if i < config.StartRow+3 {
            fmt.Printf("Row %d cleaned: Date=%s, Description=%s, Amount=%s\n",
                i, date, description, amount)
		} */

		// Skip any rows where amount isn't a valid number
		_, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Printf("Skipping row %d due to invalid amount: %s\n", i, amount)
			continue
		}

		cleanedRow := []string{date, description, amount}
		cleaned = append(cleaned, cleanedRow)
	}
	return cleaned
}

func processOneFile(filename string, config CSVConfig) ([][]string, error) {
    records, err := ReadCSVFile(filename, config)
    if err != nil {
        return nil, fmt.Errorf("error processing %s: %v", filename, err)
    }
    
    cleanedRecords := cleanTransactions(records, config)
    return cleanedRecords, nil
}

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


func filterByMonth(cleaned [][]string, months []int) [][]string {
	// neue datenbank zum füllen
	var transactionBase [][]string

	for _, cleanedRow := range cleaned {

		// split und umformatieren
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
