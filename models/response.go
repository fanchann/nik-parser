package models

type Response struct {
	NIK        string `json:"nik"`
	Gender     string `json:"gender"`
	Province   string `json:"province"`
	City       string `json:"city"`
	District   string `json:"district"`
	Uniqcode   string `json:"uniqcode"`
	BornDay    int    `json:"born_day"`
	BornMonth  string `json:"born_month"`
	BornYear   int    `json:"born_year"`
	PostalCode string `json:"postal_code"`
}
