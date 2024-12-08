package repository

import "github.com/google/uuid"

type ReceiptItem struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type Receipt struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ReceiptItem `json:"items"`
	Total        float64       `json:"total"`
}

type ReceiptRepository struct {
	db map[string]Receipt
}

func NewReceiptRepository() *ReceiptRepository {
	return &ReceiptRepository{db: make(map[string]Receipt)}
}

func (r *ReceiptRepository) Create(receipt Receipt) string {
	id := uuid.New().String()
	r.db[id] = receipt
	return id
}

func (r *ReceiptRepository) GetByID(id string) (*Receipt, bool) {
	receipt, exists := r.db[id]
	return &receipt, exists
}
