func PrintCategoryBreakdown(totals map[string]float64) {
    // Define your desired category order
    orderedCategories := []string{
        "rent",
        "flat",
        "food",
        // Add all your categories in the order you want them
    }

    fmt.Println("\nSpending by Category")
    fmt.Println("-----------------")
    
    // Print categories in the defined order
    for _, category := range orderedCategories {
        if amount, exists := totals[category]; exists {
            fmt.Printf("%-15s %10.2f kr\n", category+":", amount)
        }
    }

    // Optionally: print any categories that weren't in your ordered list
    for category, amount := range totals {
        if !contains(orderedCategories, category) {
            fmt.Printf("%-15s %10.2f kr\n", category+":", amount)
        }
    }
}

// Helper function to check if a string is in a slice
func contains(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}