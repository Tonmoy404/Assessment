package service

type Product struct {
	Id            string   `json:"Id"`
	Name          string   `json:"Name"`
	Description   string   `json:"Description"`
	Specification string   `json:"Specifications"`
	BrandId       string   `json:"BrandId"`
	CategoryId    string   `json:"CategoryId"`
	SupplierId    string   `json:"SupplierId"`
	UnitPrice     float64  `json:"UnitPrice"`
	DiscountPrice float64  `json:"DiscountPrice"`
	Tags          []string `json:"Tags"`
	StatusId      string   `json:"StatusId"`
	CreatedAt     int64    `json:"CreatedAt"`
}

type FilterProducts struct {
	Name       string  `json:"Name"`
	MaxPrice   float64 `json:"MaxPrice"`
	MinPrice   float64 `json:"MinPrice"`
	BrandId    string  `json:"BrandId"`
	CategoryId string  `json:"CategoryId"`
	SupplierId string  `json:"SupplierId"`
	Limit      int64   `json:"Limit"`
}
