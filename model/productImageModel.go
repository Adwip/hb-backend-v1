package model

type ProductImagesResponse struct {
	ImageName string `json:"fileName"`
	Base64    string `json:"imageFile"`
	MainImage bool   `json:"mainImage"`
}
