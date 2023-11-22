package rest

type createBrandReq struct {
	Name     string `json:"name" binding:"required"`
	StatusId int64  `json:"status_id" binding:"required"`
}

type getBrandsReq struct {
	Page  int64 `json:"page" binding:"required,min=1"`
	Limit int64 `json:"limit" binding:"required,min=1"`
}

type updateBrandReq struct {
	Name     string `json:"name" binding:"required,min=2"`
	StatusID int64    `json:"status_id" binding:"required"`
}