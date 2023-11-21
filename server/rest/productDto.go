package rest

type createProductReq struct {
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	Specification string   `json:"specifications" binding:"required"`
	BrandId       string   `json:"brand_id" binding:"required"`
	CategoryId    string   `json:"vategory_id" binding:"required"`
	SupplierId    string   `json:"supplierId" binding:"required"`
	UnitPrice     float64  `json:"unit_price" binding:"required"`
	DiscountPrice float64  `json:"discount_price" binding:"required"`
	Tags          []string `json:"tags" binding:"required"`
	StatusId      string   `json:"status_id" inding:"required"`
	CreatedAt     int64    `json:"created_at" binding:"required"`
}

type filterProductsReq struct {
	Name       string  `json:"name"`
	MaxPrice   float64 `json:"max_price"`
	MinPrice   float64 `json:"minPrice"`
	BrandId    string  `json:"brand_id"`
	CategoryId string  `json:"category_id"`
	SupplierId string  `json:"supplier_id"`
	Limit      int64   `json:"limit"`
}
