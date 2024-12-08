package service

import (
	"testing"

	"github.com/IslaMurtazaev/receipt-processor/repository"
	"github.com/stretchr/testify/assert"
)

func TestCountRetailerBonus(t *testing.T) {
	service := NewReceiptPointsService()

	tests := []struct {
		name           string
		retailerName   string
		expectedPoints int
	}{
		{"Empty retailer name", "", 0},
		{"Retailer name with letters", "BestStore", 9},
		{"Retailer name with digits", "Store123", 8},
		{"Retailer name with special characters", "Store!@#", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.countRetailerBonus(tt.retailerName)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}

func TestCountTotalFieldBonus(t *testing.T) {
	service := NewReceiptPointsService()

	tests := []struct {
		name           string
		receiptTotal   float64
		expectedPoints int
	}{
		{"Exact dollar total + multiple of 0.25", 100.00, 75},
		{"Multiple of 0.25", 50.25, 25},
		{"Other value", 99.99, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.countTotalFieldBonus(tt.receiptTotal)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}

func TestCountItemsBonus(t *testing.T) {
	service := NewReceiptPointsService()

	tests := []struct {
		name           string
		receiptItems   []repository.ReceiptItem
		expectedPoints int
	}{
		{"No items", []repository.ReceiptItem{}, 0},
		{"Two items", []repository.ReceiptItem{
			{ShortDescription: "a", Price: 5.0},
			{ShortDescription: "b", Price: 10.0},
		}, 5},
		{"Four items with matching descriptions", []repository.ReceiptItem{
			{ShortDescription: "abc", Price: 5.0},
			{ShortDescription: "xyz", Price: 10.0},
			{ShortDescription: "ghi", Price: 15.0},
			{ShortDescription: "xyz", Price: 21.0},
		}, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.countItemsBonus(tt.receiptItems)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}

func TestCountDateBonus(t *testing.T) {
	service := NewReceiptPointsService()

	tests := []struct {
		name           string
		purchaseDate   string
		expectedPoints int
	}{
		{"Odd day", "2024-12-01", 6},
		{"Even day", "2024-12-02", 0},
		{"Another odd day", "2024-12-03", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.countDateBonus(tt.purchaseDate)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}

func TestCountTimeBonus(t *testing.T) {
	service := NewReceiptPointsService()

	tests := []struct {
		name           string
		purchaseTime   string
		expectedPoints int
	}{
		{"Before 2pm", "13:30", 0},
		{"Between 2pm and 4pm", "15:00", 10},
		{"After 4pm", "16:00", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.countTimeBonus(tt.purchaseTime)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}

func TestCalculate(t *testing.T) {
	receipt1 := repository.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []repository.ReceiptItem{
			{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
			{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
			{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
			{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
		},
		Total: 35.35,
	}

	receipt2 := repository.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []repository.ReceiptItem{
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
		},
		Total: 9.00,
	}

	service := NewReceiptPointsService()

	points1 := service.Calculate(receipt1)
	expectedPoints1 := 28
	assert.Equal(t, expectedPoints1, points1, "Calculated points for receipt1 should match the expected value.")

	points2 := service.Calculate(receipt2)
	expectedPoints2 := 109
	assert.Equal(t, expectedPoints2, points2, "Calculated points for receipt2 should match the expected value.")
}
