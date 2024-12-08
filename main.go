package main

import (
	"fmt"
	"github.com/IslaMurtazaev/receipt-processor/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	receiptRepository := repository.NewReceiptRepository()

	r.POST("/receipts/process", func(c *gin.Context) {
		var receipt repository.Receipt

		c.ShouldBindJSON(&receipt)

		fmt.Printf("%+v\n", receipt)

		receiptId := receiptRepository.Create(receipt)

		c.JSON(http.StatusOK, receiptId)
	})

	r.GET("/receipts/:id/points", func(c *gin.Context) {
		var receiptId = c.Param("id")

		fmt.Printf("%+v\n", receiptId)

		receipt := receiptRepository.GetByID(receiptId)

		c.JSON(http.StatusOK, receipt)
	})

	r.Run(":8080")
}
