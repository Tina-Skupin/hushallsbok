# Financial Report Generator

A command-line tool that processes financial data and generates comprehensive financial reports in both text and CSV formats.

## Features

- Analyzes financial transactions from CSV data files
- Creates detailed financial reports with category breakdowns
- Outputs reports in both text and CSV formats
- Supports filtering by specific months and years
- Categorizes transactions automatically

## Installation

1. Ensure you have Go installed on your system
2. Clone this repository
3. Build the project with `go build`

## Usage

### Command-line arguments:
# add your individual information in main.go:

## Input File Format
The input CSV should contain the following columns:
- Date (DD.MM.YYYY format)
- Description
- Amount (using point as decimal separator)

the program takes and prepares the infofile from your dowloaded banc file. you need to add in where the relevant information is in 
"config"
- StartRow: if your infofile has header information you tell the code to skip it here. Enter the first line with transactions
- DateCol:  field that contains transcation date
- DescriptionCol: decription field
- AmountCol: field that contains the expense

- `-year`: The year to analyze (required)
- `-months`: Comma-separated list of months to analyze (required)
- `-file`: Path to the input CSV file (optional, defaults to "financials.csv")

## Output

The program generates two files:
- `financial_reportYYYY_MM.txt`: Text report with detailed breakdown
- `financial_reportYYYY_MM.csv`: CSV report for spreadsheet analysis