package service

type ProductStock struct {
	Id            string `json:"Id"`
	ProductId     string `json:"ProductId"`
	StockQuantity int64  `json:"StockQuantity"`
	UpdatedAt     int64  `json:"UpdatedAt"`
}
