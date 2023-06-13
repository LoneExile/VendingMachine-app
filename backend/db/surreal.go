package db

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"vending/config"

	"github.com/surrealdb/surrealdb.go"
)

type DBClient interface {
	InitDB(cfg *config.Config) (*dbClient, error)
	GetByTableName(tableName string) (interface{}, error)
	PostCheckout(cart Cart, pocket Pocket) (interface{}, error)
}

type dbClient struct {
	db *surrealdb.DB
}

func NewDB() *dbClient {
	return &dbClient{}
}

func (dbClient *dbClient) InitDB(cfg *config.Config) (*dbClient, error) {
	db, err := NewSurrealDBClient(cfg)
	if err != nil {
		return nil, err
	}

	dbClient.db = db

	return dbClient, nil
}

func NewSurrealDBClient(cfg *config.Config) (*surrealdb.DB, error) {
	db, err := surrealdb.New(cfg.DBEndpoint)
	if err != nil {
		return nil, err
	}
	log.Println("🟢 Connected to SurrealDB")

	// Sign in
	if _, err = db.Signin(map[string]string{
		"user": cfg.DBUser,
		"pass": cfg.DBPassword,
	}); err != nil {
		log.Println("🔴 Error signing in to SurrealDB:", err)
		return nil, err
	}
	log.Println("🟢 Signed in to SurrealDB")

	// Select namespace and database
	if _, err = db.Use(cfg.DBNamespace, cfg.DBDatabase); err != nil {
		return nil, err
	}
	log.Println("🟢 Selected namespace:", cfg.DBNamespace)
	log.Println("🟢 Selected database:", cfg.DBDatabase)

	return db, nil
}

func (dbClient *dbClient) GetByTableName(tableName string) (interface{}, error) {
	result, err := dbClient.db.Select(tableName)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateStock(cart Cart, dbClient *dbClient, returnMessage map[string]interface{}) (map[string]interface{}, error) {
	// update stock product
	for _, item := range cart.Items {
		productResult, err := dbClient.db.Select(item.ID)
		if err != nil {
			returnMessage["status"] = "failed"
			returnMessage["message"] = "Product not found"
			return returnMessage, nil
		}

		product := productResult.(map[string]interface{})
		product["stock"] = product["stock"].(float64) - float64(item.Quantity)

		_, err = dbClient.db.Update(item.ID, product)
		if err != nil {
			returnMessage["status"] = "failed"
			returnMessage["message"] = "failed to update product"
			return returnMessage, nil
		}
	}
	return returnMessage, nil
}

func (dbClient *dbClient) PostCheckout(cart Cart, pocket Pocket) (interface{}, error) {
	returnMessage := make(map[string]interface{})
	totalCost := cart.Total
	log.Println("🟢 Total cost:", totalCost)
	totalChange := pocket.Balance - totalCost
	log.Println("🟢 Change:", totalChange)

	// check if total change is sufficient
	for _, item := range cart.Items {
		productResult, err := dbClient.db.Select(item.ID)
		if err != nil {
			returnMessage["status"] = "failed"
			returnMessage["message"] = "Product not found"
			return returnMessage, nil
		}

		product := productResult.(map[string]interface{})
		if product["stock"].(float64) < float64(item.Quantity) {
			returnMessage["status"] = "failed"
			returnMessage["message"] = fmt.Sprintf("Insufficient stock for product %s", item.ID)
			return returnMessage, nil
		}
	}

	if totalChange < 0 {
		returnMessage["message"] = "Insufficient balance"
		return returnMessage, nil
	} else if totalChange == 0 {
		for _, item := range pocket.Items {
			denominationResult, err := dbClient.db.Select(item.ID)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "Denomination not found"
				return returnMessage, nil
			}

			denomination := denominationResult.(map[string]interface{})
			denomination["stock"] = denomination["stock"].(float64) + float64(item.Quantity)

			_, err = dbClient.db.Update(item.ID, denomination)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "failed to update denomination"
				return returnMessage, nil
			}
		}

		m, err := UpdateStock(cart, dbClient, returnMessage)
		if err != nil {
			return m, err
		}

		returnMessage["status"] = "success"
		returnMessage["message"] = "No change"
		return returnMessage, nil
	} else {
		// descending
		sort.Slice(pocket.Items, func(i, j int) bool {
			iValue, err := strconv.ParseFloat(pocket.Items[i].DenominationValue, 64)
			if err != nil {
				log.Fatal(err)
			}

			jValue, err := strconv.ParseFloat(pocket.Items[j].DenominationValue, 64)
			if err != nil {
				log.Fatal(err)
			}

			return iValue > jValue
		})

		changeDenominations := Pocket{
			Items:   []PocketItem{},
			Balance: 0,
		}

		// calculate change
		for _, denomination := range pocket.Items {
			denominationValue, err := strconv.ParseFloat(denomination.DenominationValue, 64)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "failed to parse denomination value"
				return returnMessage, nil
			}
			maxDenomination := int(totalChange / denominationValue)

			if denomination.Stock < maxDenomination {
				maxDenomination = denomination.Stock
			}

			totalChange -= float64(maxDenomination) * denominationValue

			if maxDenomination > 0 {
				changeDenominations.Items = append(changeDenominations.Items, PocketItem{
					ID:                denomination.ID,
					DenominationValue: denomination.DenominationValue,
					Quantity:          maxDenomination,
					Typed:             denomination.Typed,
				})
			}

			if totalChange == 0 {
				break
			}
		}

		for _, logChange := range changeDenominations.Items {
			log.Println("🟢 Change:", logChange.Quantity, "ea ", logChange.DenominationValue, " THB")
		}

		// add to db before removing
		for _, item := range pocket.Items {
			denominationResult, err := dbClient.db.Select(item.ID)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "Denomination not found"
				return returnMessage, nil
			}

			denomination := denominationResult.(map[string]interface{})
			denomination["stock"] = denomination["stock"].(float64) + float64(item.Quantity)

			_, err = dbClient.db.Update(item.ID, denomination)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "failed to update denomination"
				return returnMessage, nil
			}
		}

		// remove stock from db
		for _, item := range changeDenominations.Items {
			denominationResult, err := dbClient.db.Select(item.ID)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "Denomination not found"
				return returnMessage, nil
			}
			denomination := denominationResult.(map[string]interface{})
			currentStock := denomination["stock"].(float64)

			newStock := currentStock - float64(item.Quantity)
			denomination["stock"] = newStock

			_, err = dbClient.db.Update(item.ID, denomination)
			if err != nil {
				returnMessage["status"] = "failed"
				returnMessage["message"] = "failed to update denomination"
				return returnMessage, nil
			}
		}

		m, err := UpdateStock(cart, dbClient, returnMessage)
		if err != nil {
			return m, err
		}

		returnMessage["status"] = "success"
		returnMessage["message"] = "Successfull transaction"
		returnMessage["change"] = changeDenominations.Items
		log.Println("🟢 last:", changeDenominations.Items)
		return returnMessage, nil
	}
}
