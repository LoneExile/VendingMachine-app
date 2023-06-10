package db

import (
	"log"
	"vending/config"

	"github.com/surrealdb/surrealdb.go"
)

func Init(cfg *config.Config) (*surrealdb.DB, error) {
	db, err := surrealdb.New(cfg.DBEndpoint)
	if err != nil {
		panic(err)
	}
	log.Println("🟢 Connected to SurrealDB")

	// Sign in
	if _, err = db.Signin(map[string]string{
		"user": cfg.DBUser,
		"pass": cfg.DBPassword,
	}); err != nil {
		panic(err)
	}
	log.Println("🟢 Signed in to SurrealDB")

	// Select namespace and database
	if _, err = db.Use(cfg.DBNamespace, cfg.DBDatabase); err != nil {
		panic(err)
	}
	log.Println("🟢 Selected namespace:", cfg.DBNamespace)
	log.Println("🟢 Selected database:", cfg.DBDatabase)

	return db, nil
}

func GetByTableName(db *surrealdb.DB, tableName string) (interface{}, error) {
	result, err := db.Select(tableName)
	if err != nil {
		return nil, err
	}
	return result, nil
}
