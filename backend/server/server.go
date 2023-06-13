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
			"denominations": result,
		})
	})

	r.POST("/checkout", func(c *gin.Context) {
		var requestBody struct {
			Cart   db.Cart   `json:"cart"`
			Pocket db.Pocket `json:"pocket"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		result, err := apiInstance.Checkout(requestBody.Cart, requestBody.Pocket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", cfg.NextPublicServerIP + ":3000", "https://bpi.voidbox.dev", "http://bpi.voidbox.dev"},
		AllowCredentials: true,
	})

	http.ListenAndServe(":8080", corsMiddleware.Handler(r))

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
