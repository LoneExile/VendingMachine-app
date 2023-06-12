package server

import (
	"net/http"
	"vending/api"
	"vending/config"
	"vending/db"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func StartServer(cfg *config.Config, dbClient db.DBClient) {
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
	// GetDenomination
	r.GET("/denomination", func(c *gin.Context) {
		result, err := apiInstance.GetDenomination()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"denomination": result,
		})
	})

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", cfg.NextPublicServerIP + ":3000"},
		AllowCredentials: true,
	})

	http.ListenAndServe(":8080", corsMiddleware.Handler(r))

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
