package model

import "github.com/google/uuid"

type AllProductsResponse struct {
	ID           uuid.UUID `json:"id"`
	ProductName  string    `json:"productName"`
	Creator      string    `json:"creatore"`
	Negotiable   bool      `json:"negotiable"`
	PurchaseType string    `json:"purchaseType"`
	Favourite    int32     `json:"favourite"`
	Price        int64     `json:"price"`
	ProductImage string    `json:"image"`
}

type ProductByIDResponse struct {
	ID           uuid.UUID `json:"id"`
	ProductName  string    `json:"productName"`
	Creator      string    `json:"creatore"`
	Negotiable   bool      `json:"negotiable"`
	PurchaseType string    `json:"purchaseType"`
	Favourite    int32     `json:"favourite"`
	Price        int64     `json:"price"`
	ProductImage string    `json:"image"`
	OfferStatus  string    `json:"offerStatus"`
	Status       string    `json:"status"`
}
