package main

import (
    "testing"
)

func TestProcessOneFile(t *testing.T) {
    // Set up a test config that matches your CSV structure
    config := CSVConfig{
        StartRow:       1,  // adjust these to match your actual CSV
        DateCol:        0,
        DescriptionCol: 1,
        AmountCol:      2,
    }
    
    // Test with one of your actual CSV files
    records, err := processOneFile("2024_martin.csv", config)
    if err != nil {
        t.Fatalf("processOneFile failed: %v", err)
    }
    
    // Basic checks
    if len(records) == 0 {
        t.Error("Expected records, got empty slice")
    }
    
    // Check structure of first record
    if len(records[0]) != 3 { // expecting [date, description, amount]
        t.Errorf("Expected 3 columns, got %d", len(records[0]))
    }
}

func TestCombineTransactions(t *testing.T) {
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
        "24_12_tina.csv",   // Use your actual file names
        "2024_martin.csv",    // Use your actual file names
    }
    
	
    // Get records from individual files first
    records1, err := processOneFile(files[0], configs[0])
    if err != nil {
        t.Fatalf("Failed to process first file: %v", err)
    }
    records2, err := processOneFile(files[1], configs[1])
    if err != nil {
        t.Fatalf("Failed to process second file: %v", err)
    }

    // Now get combined records
    combined, err := combineTransactions(files, configs)
    if err != nil {
        t.Fatalf("combineTransactions failed: %v", err)
    }

    // Verify the counts
    expectedTotal := len(records1) + len(records2)
    if len(combined) != expectedTotal {
        t.Errorf("Expected %d total records, but got %d", expectedTotal, len(combined))
    }

    t.Logf("File 1 records: %d", len(records1))
    t.Logf("File 2 records: %d", len(records2))
    t.Logf("Combined records: %d", len(combined))
}