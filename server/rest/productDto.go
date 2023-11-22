package rest

type createProductReq struct {
	Name          string   `json:"name" binding:"required, min=2"`
	Description   string   `json:"description" binding:"required"`
	Specification string   `json:"specifications" binding:"required"`
	BrandId       string   `json:"brand_id" binding:"required"`
	CategoryId    string   `json:"vategory_id" binding:"required"`
	SupplierId    string   `json:"supplierId" binding:"required"`
	UnitPrice     float64  `json:"unit_price" binding:"required,min=0"`
	DiscountPrice float64  `json:"discount_price" binding:"required,min=0"`
	Tags          []string `json:"tags" binding:"required"`
	StatusId      int64   `json:"status_id" binding:"required"`
	Stock         int64    `json:"stock" binding:"required,min=1"`
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

type updateProductReq struct {
	Name           string   `json:"name" binding:"required,min=2"`
	Description    string   `json:"description" binding:"required"`
	Specifications string   `json:"specifications" binding:"required"`
	BrandID        string   `json:"brand_id" binding:"required"`
	CategoryID     string   `json:"category_id" binding:"required"`
	SupplierID     string   `json:"supplier_id" binding:"required"`
	UnitPrice      float64  `json:"unit_price" binding:"required,min=0"`
	DiscountPrice  float64  `json:"discount_price" binding:"required,min=0"`
	Tags           []string `json:"tags" binding:"required"`
	StatusID       int64      `json:"status_id" binding:"required,validStatusID"`
	Stock          int64    `json:"stock" binding:"required,min=1"`
}
