package model

import "github.com/google/uuid"

type AddProductRequest struct {
	Title         string         `json:"productName"`
	Field         int8           `json:"fieldType"`
	OpenNegotiate bool           `json:"openNegotiate"`
	Price         int64          `json:"price"`
	PurchaseType  string         `json:"purchaseType"`
	Type          string         `json:"type"`
	Displayed     bool           `json:"displayed"`
	Kuota         int32          `json:"kuota"`
	Status        string         `json:"status"`
	Images        []ProductImage `json:"images"`
}

type ProductImage struct {
	ImageName string `json:"fileName"`
	Base64    string `json:"imageFile"`
	MainImage bool   `json:"mainImage"`
}

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
