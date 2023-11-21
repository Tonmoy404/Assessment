package service

type Category struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	ParentId  string `json:"ParentId"`
	Sequence  string `json:"Sequence"`
	StatusId  string `json:"StatusId"`
	CreatedAt int64  `json:"CreatedAt"`
}
