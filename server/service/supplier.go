package service

type Supplier struct {
	Id                 string `json:"Id"`
	Name               string `json:"Name"`
	Email              string `json:"Email"`
	Phone              string `json:"Phone"`
	StatusId           int64 `json:"StatusId"`
	IsVerifiedSupplier bool   `json:"IsVerifiedSupplier"`
	CreatedAt          int64  `json:"CreatedAt"`
}
