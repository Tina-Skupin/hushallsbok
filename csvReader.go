package main

import (
	"encoding/csv"
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
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err

	}

	return records, err
}

func cleanTransactions(records [][]string, config CSVConfig) [][]string {
	var cleaned [][]string

	// ignore the first 8 lines
	for i := config.StartRow; i < len(records); i++ {
		row := records[i]
		if len(row) == 0 {
			continue
		}

		// Clean up each field, removing extra spaces
		date := strings.TrimSpace(row[config.DateCol])
		description := strings.TrimSpace(row[config.DescriptionCol])
		amount := strings.TrimSpace(row[config.AmountCol])
		//amount := strings.TrimSpace(row[3]) original, in case something breaks

		// Skip any rows where amount isn't a valid number
		_, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			continue
		}

		cleanedRow := []string{date, description, amount}
		cleaned = append(cleaned, cleanedRow)
	}
	return cleaned
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
