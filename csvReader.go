package main

import (
	"os"
	"encoding/csv"
	"strings"
    "strconv"
)

func ReadCSVFile(filename string) ([][]string, error) {
	// open the file
	file, err := os.Open("testcsv.csv")
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

func cleanTransactions(records [][]string) [][]string {
    var cleaned [][]string
    
    // ignore the first 8 lines
    for i := 8; i < len(records); i++ {
        row := records[i]
        if len(row) == 0 {
            continue
        }

        // Clean up each field, removing extra spaces
        date := strings.TrimSpace(row[1])
        description := strings.TrimSpace(row[2])
        amount := strings.TrimSpace(row[3])
        
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