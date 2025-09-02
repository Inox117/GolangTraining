# CsvParser

## Task

Create a program that parses a CSV file and output the data, in a json format, in the console.

The program should accept the following arguments:
    `month` in format `YYYYMM`
    `history_file`, the path to the CSV file

### File specification

    Header: `date,amount,content`
    date : string format `YYYY/MM/DD`
    amount : integer
    content : string

None of the columns can be empty.
Row order is not guaranteed.

### Output specification

Format : JSON
```json
{
 "period": "2022/01",
 "total_income": 0,
 "total_expenditure": -111000,
 "transactions": [
  {
   "date": "2022/01/25",
   "amount": -100000,
   "content": "rent"
  },
  {
   "date": "2022/01/06",
   "amount": -10000,
   "content": "debit"
  },
  {
   "date": "2022/01/05",
   "amount": -1000,
   "content": "eating out"
  }
 ]
}
```
`period` is the month in format `YYYY/MM`
`total_income` is the sum of all the incomes
`total_expenditure` is the sum of all the expenditures
`transactions` is an array of transactions
`date` is the date in format `YYYY/MM/DD`
`amount` is the amount of the transaction
`content` is the content of the transaction

The data must be ordered by date in descending order.

### Constraints
The program should process only the line for the year and month specified.

### Trivia

Handle all possible errors.
An example file is provided in the `samplingData` folder.