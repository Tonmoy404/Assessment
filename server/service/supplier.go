package service

type Supplier struct {
	Id                 string `json:"Id"`
	Name               string `json:"Name"`
	Email              string `json:"Email"`
	Phone              string `json:"Phone"`
	StatusId           string `json:"StatusId"`
	IsVerifiedSupplier string `json:"IsVerifiedSupplier"`
	CreatedAt          int64  `json:"CreatedAt"`
}
