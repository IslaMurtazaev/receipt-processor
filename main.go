package main

import (
	"fmt"
	"github.com/IslaMurtazaev/receipt-processor/repository"
	"github.com/IslaMurtazaev/receipt-processor/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ReceiptItem `json:"items"`
	Total        string        `json:"total"`
}

func main() {
	r := gin.Default()

	receiptRepository := repository.NewReceiptRepository()
	receiptService := service.NewReceiptPointsService()

	r.POST("/receipts/process", func(c *gin.Context) {
		var receipt Receipt

		if err := c.ShouldBindJSON(&receipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		parsedReceipt, err := parseReceipt(receipt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		receiptId := receiptRepository.Create(parsedReceipt)
		c.JSON(http.StatusOK, receiptId)
	})

	r.GET("/receipts/:id/points", func(c *gin.Context) {
		var receiptId = c.Param("id")
		receipt, exists := receiptRepository.GetByID(receiptId)

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
			return
		}

		receiptPoints := receiptService.Calculate(*receipt)
		c.JSON(http.StatusOK, receiptPoints)
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func parseReceipt(receipt Receipt) (repository.Receipt, error) {
	parsedReceipt := repository.Receipt{
		Retailer:     receipt.Retailer,
		PurchaseDate: receipt.PurchaseDate,
		PurchaseTime: receipt.PurchaseTime,
	}

	// Convert string item prices and total to float64
	for i := range receipt.Items {
		parsedPrice, err := strconv.ParseFloat(receipt.Items[i].Price, 64)
		if err != nil {
			return parsedReceipt, err
		}
		parsedReceipt.Items = append(parsedReceipt.Items, repository.ReceiptItem{
			ShortDescription: receipt.Items[i].ShortDescription,
			Price:            parsedPrice,
		})
	}

	parsedTotal, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return parsedReceipt, err
	}
	parsedReceipt.Total = parsedTotal

	return parsedReceipt, nil
}
