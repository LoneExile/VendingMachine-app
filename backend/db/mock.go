package db

import (
	"fmt"
	"log"
)

func (dbClient *dbClient) Mock() {
	db := dbClient.db

	// Create Products
	productData := []Product{
		{ProductID: 1, ProductName: "Water1", Price: 1, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 2, ProductName: "Water2", Price: 2, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 3, ProductName: "Water3", Price: 3, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 4, ProductName: "Water4", Price: 4, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 5, ProductName: "Water5", Price: 5, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 6, ProductName: "Water6", Price: 6, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 7, ProductName: "Water7", Price: 7, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 8, ProductName: "Water8", Price: 8, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 9, ProductName: "Water9", Price: 9, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 10, ProductName: "Water10", Price: 10, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 11, ProductName: "Water11", Price: 11, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
		{ProductID: 12, ProductName: "Water12", Price: 12, Stock: 10, Picture: "https://illustoon.com/photo/11612.png"},
	}

	for _, p := range productData {
		_, err := db.Query(fmt.Sprintf(`
			INSERT INTO product {id: "product:%d", product_name: "%s", price: %.2f, stock: %d, picture: "%s"}
			`,
			p.ProductID, p.ProductName, p.Price, p.Stock, p.Picture),
			nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create Coins and Banknotes
	denominationData := []CoinAndBanknote{
		{DenominationID: 1, DenominationValue: 1, Stock: 100, Typed: "coin"},
		{DenominationID: 2, DenominationValue: 5, Stock: 100, Typed: "coin"},
		{DenominationID: 3, DenominationValue: 10, Stock: 100, Typed: "coin"},
		{DenominationID: 4, DenominationValue: 20, Stock: 50, Typed: "banknote"},
		{DenominationID: 5, DenominationValue: 50, Stock: 50, Typed: "banknote"},
		{DenominationID: 6, DenominationValue: 100, Stock: 50, Typed: "banknote"},
		{DenominationID: 7, DenominationValue: 500, Stock: 10, Typed: "banknote"},
		{DenominationID: 8, DenominationValue: 1000, Stock: 5, Typed: "banknote"},
	}

	for _, d := range denominationData {
		_, err := db.Query(fmt.Sprintf(`INSERT INTO denomination {id: "denomination:%d", denomination_value: %.2f, stock: %d, typed: "%s"}`,
			d.DenominationID, d.DenominationValue, d.Stock, d.Typed),
			nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create Transactions
	transactionData := []Transaction{
		{TransactionID: 1, ProductID: 1, Quantity: 1, TotalPrice: 1.5, TransactionTime: "2021-09-01 12:00:00"},
		{TransactionID: 2, ProductID: 2, Quantity: 1, TotalPrice: 1.5, TransactionTime: "2021-09-01 12:00:00"},
		{TransactionID: 3, ProductID: 3, Quantity: 1, TotalPrice: 1.0, TransactionTime: "2021-09-01 12:00:00"},
		{TransactionID: 4, ProductID: 1, Quantity: 1, TotalPrice: 1.5, TransactionTime: "2021-09-01 12:00:00"},
		{TransactionID: 5, ProductID: 2, Quantity: 1, TotalPrice: 1.5, TransactionTime: "2021-09-01 12:00:00"},
	}

	for _, t := range transactionData {
		_, err := db.Query(fmt.Sprintf(`INSERT INTO transaction {id: "transaction:%d", product_id: product:%d, quantity: %d, total_price: %.2f, transaction_time: "%s"}`, t.TransactionID, t.ProductID, t.Quantity, t.TotalPrice, t.TransactionTime), nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	paymentData := []Payment{
		{PaymentID: 1, TransactionID: 1, DenominationID: 1, Quantity: 2, TotalValue: 2.0},
		{PaymentID: 2, TransactionID: 2, DenominationID: 2, Quantity: 1, TotalValue: 5.0},
		{PaymentID: 3, TransactionID: 3, DenominationID: 3, Quantity: 1, TotalValue: 10.0},
		{PaymentID: 4, TransactionID: 4, DenominationID: 4, Quantity: 1, TotalValue: 20.0},
		{PaymentID: 5, TransactionID: 5, DenominationID: 5, Quantity: 1, TotalValue: 50.0},
	}

	for _, p := range paymentData {
		_, err := db.Query(fmt.Sprintf(`INSERT INTO payment {id: "payment:%d", transaction_id: "transaction:%d", denomination_id: "denomination:%d", quantity: %d, total_value: %.2f}`, p.PaymentID, p.TransactionID, p.DenominationID, p.Quantity, p.TotalValue), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("🟢 Mocked database")
}
