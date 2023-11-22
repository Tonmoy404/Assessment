package service

type Brand struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	StatusId  int64  `json:"StatusId"`
	CreatedAt int64  `json:"CreatedAt"`
}

type BrandResult struct {
	Brands []*Brand `json:"brands"`
	Total  int      `json:"total"`
	Page   int64    `json:"page"`
	Limit  int64    `json:"limit"`
}
