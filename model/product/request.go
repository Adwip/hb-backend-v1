package product

type AddProduct struct {
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

type ModifyProductImages struct {
	Product      string                `json:"product"`
	NewImages    []ProductImage        `json:"newImages"`
	ModifiedFile []DeletedProductImage `json:"modifiedImages"`
}

type ProductImage struct {
	ImageName string `json:"fileName"`
	Base64    string `json:"imageFile"`
	MainImage bool   `json:"mainImage"`
}

type DeletedProductImage struct {
	ImageID string `json:"imageID"`
	Action  int8   `json:"action"`
}
