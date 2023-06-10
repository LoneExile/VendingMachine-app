package api

import (
	"vending/config"
	"vending/db"

	"github.com/surrealdb/surrealdb.go"
)

type ginAPI struct {
	dbClient *surrealdb.DB
}

func NewAPI(cfg *config.Config, dbClient *surrealdb.DB) (*ginAPI, error) {
	return &ginAPI{dbClient: dbClient}, nil
}

func (api *ginAPI) GetProducts() (interface{}, error) {
	result, err := db.GetByTableName(api.dbClient, "product")
	if err != nil {
		return nil, err
	}
	return result, nil
}
