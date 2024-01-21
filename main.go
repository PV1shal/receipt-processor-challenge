package main

import (
	"fmt"
	"receipt-processor-challenge/pkg"
	"receipt-processor-challenge/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Inintializing server...")

	reciptStore := pkg.NewMemReciptStore()                // In-Memory Storage
	reciptHandler := routes.NewReciptHandler(reciptStore) // Recipt Handler

	router := gin.Default()

	router.POST("/receipts/process", reciptHandler.AddNewRecipt)
	router.GET("/receipts/:ID/points", reciptHandler.GetRecipt)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	router.Run("localhost:8080")
}

func pingTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
