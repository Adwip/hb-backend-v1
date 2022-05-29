package product

type ProductByIdResponse struct {
	ID           string
	Name         string
	Image        string
	PurchaseType int8
	IsNegotiable bool
	Field        int8
	Images       []ProductImage
}
