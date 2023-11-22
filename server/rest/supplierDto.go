package rest

type createSupplierReq struct {
	Name               string `json:"name" binding:"required,min=2"`
	Email              string `json:"email" binding:"required,email"`
	Phone              string `json:"phone" binding:"required,validPhone"`
	StatusId           int64  `json:"status_id" binding:"required"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" binding:"required"`
}

type getSuppliersReq struct {
	Page  int64 `json:"page" binding:"required,min=1"`
	Limit int64 `json:"limit" binding:"required,min=1"`
}

type updateSupplierReq struct {
	Name               string `json:"name" binding:"required,min=2"`
	Email              string `json:"email" binding:"required,email"`
	Phone              string `json:"phone" binding:"required"`
	StatusID           int    `json:"status_id" binding:"required"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" binding:"required"`
}
