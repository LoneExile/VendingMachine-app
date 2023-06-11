package server

import (
	"net/http"
	"vending/api"
	"vending/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/surrealdb/surrealdb.go"
)

func StartServer(cfg *config.Config, dbClient *surrealdb.DB) {
	r := gin.Default()
	apiInstance, err := api.NewAPI(cfg, dbClient)
	if err != nil {
		panic(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products", func(c *gin.Context) {
		result, err := apiInstance.GetProducts()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"products": result,
		})
	})

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", cfg.NextPublicServerIP + ":3000"},
		AllowCredentials: true,
	})

	http.ListenAndServe(":8080", corsMiddleware.Handler(r))

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
