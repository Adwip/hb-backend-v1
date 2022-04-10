package product

type AddProduct struct {
	Title     string         `json:"product_name"`
	Field     int8           `json:"fieldType"`
	Negotiate bool           `json:"isNegotiate"`
	Images    []ProductImage `json:"images"`
	Harga     int64          `json:"price"`
	SalesType int8           `json:"salesType"`
}

type ModifyProductImages struct {
	Product     string                `json:"product"`
	NewImages   []ProductImage        `json:"newImages"`
	DeletedFile []DeletedProductImage `json:"deletedImages"`
}

type ProductImage struct {
	Base64    string `json:"imageFile"`
	MainImage bool   `json:"mainImage"`
}

type DeletedProductImage struct {
	ImageID string `json:"imageID"`
}
