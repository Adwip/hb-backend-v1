package product

type AddProduct struct {
	User            string `json:"user"`
	Jenis_produk    int8   `json:"type"`
	Kategori_produk int8   `json:"category"`
	Judul           string `json:"title"`
	Negosiasi       int8   `json:"negotiate"`
}
