package main // Use your actual package name

// FinancialSummary contains all calculated financial data from transactions
type FinancialSummary struct {
    // Transaction counts
    TotalTransactions   int
    MatchedTransactions int
    
    // Financial totals
    TotalIncome      float64
    TotalCosts       float64
    
    // Categorization information
    TotalsByCategory      map[string]float64
    TotalAllCategorized    float64
    
    // Time period
    Year   int
    Months []int
}


// FinancialSummary struct definition
    // Your current variable       // New struct field
    // ------------------          // ---------------
    // totalTransactions     ->    TotalTransactions   int
    // matched               ->    MatchedTransactions int
    // totalIncome           ->    TotalIncome         float64
    // totalCosts            ->    TotalCosts          float64
    // categoryTotals        ->    TotalsByCategory      map[string]float64
    // matchies (or matchedSum) -> TotalAllCategorized    float64
    // year                  ->    Year                int
    // months                ->    Months              []int

// Variable mapping guide:
// Old variable name -> Struct field
// totalTransactions -> summary.TotalTransactions
// matched -> summary.MatchedTransactions
// totalIncome -> summary.TotalIncome
// totalCosts -> summary.TotalCosts
// categoryTotals -> summary.TotalsByCategory
// matchies/matchedSum -> summary.TotalAllCategorized


// i put my other notes here as well...



// Calculator:
//TODO  variable categories should be called "transactionsCat"

//contains 2 functions: 
// Categorize expenses: puts all the transactions into categories using the transactionsCat map
// TODO function calculateQualityIncomeCosts should be called calculateFinances
// categorizeExpenses, that should return the sums by category map and the total of all categorized items (summary.TotalsByCategory and summary.TotalAllCategorized) 

// add a thrid function calculateFinances that calls the other two
// and fills the struct we have.