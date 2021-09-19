package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type DegiroTransaction struct {
	Customer                 string
	Date                     string
	Time                     string
	Product                  string
	ISIN                     string
	Market                   string
	ExecutionCenter          string
	Amount                   int
	Price                    float64
	PriceCurrency            string
	LocalValue               float64
	LocalValueCurrency       string
	Value                    float64
	ValueCurrency            string
	ExchangeRate             float64
	TransactionCosts         float64
	TransactionCostsCurrency string
	Total                    float64
	TotalCurrency            string
	TransactionId            string
}

func ParseDegiroTransactionsFile(filename string) (transactions []DegiroTransaction) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// var transactions []Transaction
	for index, line := range csvLines {
		if index == 0 { // remove header
			continue
		}

		amount, _ := strconv.Atoi(line[6])
		price, _ := strconv.ParseFloat(line[7], 32)
		localValue, _ := strconv.ParseFloat(line[9], 32)
		value, _ := strconv.ParseFloat(line[11], 32)
		exchangeRate, _ := strconv.ParseFloat(line[13], 32)
		transactionCosts, _ := strconv.ParseFloat(line[14], 32)
		total, _ := strconv.ParseFloat(line[16], 32)
		transaction := DegiroTransaction{
			Customer:                 "Wuerbo",
			Date:                     line[0],
			Time:                     line[1],
			Product:                  line[2],
			ISIN:                     line[3],
			Market:                   line[4],
			ExecutionCenter:          line[5],
			Amount:                   amount,
			Price:                    price,
			PriceCurrency:            line[8],
			LocalValue:               localValue,
			LocalValueCurrency:       line[10],
			Value:                    value,
			ValueCurrency:            line[12],
			ExchangeRate:             exchangeRate,
			TransactionCosts:         transactionCosts,
			TransactionCostsCurrency: line[15],
			Total:                    total,
			TotalCurrency:            line[17],
			TransactionId:            line[18],
		}
		// fmt.Println(transaction)
		transactions = append(transactions, transaction)
	}

	return
}
