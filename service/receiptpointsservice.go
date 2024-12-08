package service

import (
	"github.com/IslaMurtazaev/receipt-processor/repository"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type ReceiptPointsService struct{}

func NewReceiptPointsService() *ReceiptPointsService {
	return &ReceiptPointsService{}
}

func (s *ReceiptPointsService) countRetailerBonus(retailerName string) int {
	points := 0
	// One point for every alphanumeric character in the retailer name.
	for _, char := range retailerName {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	return points
}

func isInteger(f float64) bool {
	return f == math.Trunc(f)
}

func (s *ReceiptPointsService) countTotalFieldBonus(receiptTotal float64) int {
	points := 0

	// 50 points if the total is a round dollar amount with no cents
	if isInteger(receiptTotal) {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25
	if math.Mod(receiptTotal, 0.25) == 0 {
		points += 25
	}

	return points
}

func (s *ReceiptPointsService) countItemsBonus(receiptItems []repository.ReceiptItem) int {
	points := 0

	// 5 points for every two items on the receipt
	points += len(receiptItems) / 2 * 5

	// if trimmed description length is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer
	for _, item := range receiptItems {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	return points
}

func (s *ReceiptPointsService) countDateBonus(purchaseDate string) int {
	points := 0

	// 6 points if the day in the purchase date is odd
	dateParts := strings.Split(purchaseDate, "-")
	day, _ := strconv.ParseInt(dateParts[2], 10, 8)
	if day%2 != 0 {
		points += 6
	}

	return points
}

func (s *ReceiptPointsService) countTimeBonus(purchaseTime string) int {
	points := 0

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	timeParts := strings.Split(purchaseTime, ":")
	hour, _ := strconv.ParseInt(timeParts[0], 10, 8)
	if hour >= 14 && hour < 16 {
		points += 10
	}

	return points
}

func (s *ReceiptPointsService) Calculate(receipt repository.Receipt) int {
	var totalPoints int

	totalPoints += s.countRetailerBonus(receipt.Retailer)
	totalPoints += s.countTotalFieldBonus(receipt.Total)
	totalPoints += s.countItemsBonus(receipt.Items)
	totalPoints += s.countDateBonus(receipt.PurchaseDate)
	totalPoints += s.countTimeBonus(receipt.PurchaseTime)

	return totalPoints
}
