package db

type Product struct {
	ProductID    int           `json:"product_id" surrealdb:"pk"`
	ProductName  string        `json:"product_name"`
	Price        float64       `json:"price"`
	Stock        int           `json:"stock"`
	Picture      string        `json:"picture"`
	Transactions []Transaction `json:"transactions" surrealdb:"fk=ProductID"`
}

type Transaction struct {
	TransactionID   int       `json:"transaction_id" surrealdb:"pk"`
	ProductID       int       `json:"product_id" surrealdb:"fk"`
	Quantity        int       `json:"quantity"`
	TotalPrice      float64   `json:"total_price"`
	TransactionTime string    `json:"transaction_time"`
	Payments        []Payment `json:"payments" surrealdb:"fk=TransactionID"`
}

type CoinAndBanknote struct {
	DenominationID    int       `json:"denomination_id" surrealdb:"pk"`
	DenominationValue float64   `json:"denomination_value"`
	Stock             int       `json:"stock"`
	Typed             string    `json:"typed"`
	Payments          []Payment `json:"payments" surrealdb:"fk=DenominationID"`
}

type Payment struct {
	PaymentID      int     `json:"payment_id" surrealdb:"pk"`
	TransactionID  int     `json:"transaction_id" surrealdb:"fk"`
	DenominationID int     `json:"denomination_id" surrealdb:"fk"`
	Quantity       int     `json:"quantity"`
	TotalValue     float64 `json:"total_value"`
}
