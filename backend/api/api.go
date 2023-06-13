package api

import (
	"vending/config"
	"vending/db"
)

type ginAPI struct {
	dbClient db.DBClient
}

func NewAPI(cfg *config.Config, dbClient db.DBClient) (*ginAPI, error) {
	return &ginAPI{
		dbClient: dbClient,
	}, nil
}

func (api *ginAPI) GetProducts() (interface{}, error) {
	result, err := api.dbClient.GetByTableName("product")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api *ginAPI) GetDenomination() (interface{}, error) {
	result, err := api.dbClient.GetByTableName("denomination")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api *ginAPI) Checkout(cart db.Cart, pocket db.Pocket) (interface{}, error) {
	result, err := api.dbClient.PostCheckout(cart, pocket)
	if err != nil {
		return nil, err
	}

	return result, nil
}
