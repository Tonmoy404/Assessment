package service

type Brand struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	StatusId  string `json:"StatusId"`
	CreatedAt int64  `json:"CreatedAt"`
}
