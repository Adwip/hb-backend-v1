package model

type AllProductsResponse struct {
	ID           string `json:"id"`
	ProductName  string `json:"productName"`
	Creator      string `json:"creatore"`
	Negotiable   bool   `json:"negotiable"`
	PurchaseType string `json:"purchaseType"`
	Favourite    int32  `json:"favourite"`
	Price        int64  `json:"price"`
	ProductImage string `json:"image"`
}
