package main

import (
	"log"
	"vending/config"
	"vending/db"
	"vending/server"
	"vending/storage"
)

func main() {
	cfg := config.LoadConfig()

	store := storage.NewMinioStorage()
	client, err := store.InitClient(cfg)
	if err != nil {
		log.Println("🔴 Error initializing storage client:", err)
		return
	}

	err = client.CreateBucket()
	if err != nil {
		log.Println("🔴 Error creating bucket:", err)
		return
	}

	dbClient, err := db.Init(cfg)
	if err != nil {
		log.Println("🔴 Error initializing database client:", err)
		return
	}

	db.Mock(dbClient)

	server.StartServer(cfg, dbClient)
}
