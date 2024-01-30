package models

type Response struct {
	NIK        string `json:"nik"`
	Gender     string `json:"kelamin"`
	Province   string `json:"provinsi"`
	City       string `json:"kotakab"`
	District   string `json:"kecamatan"`
	Uniqcode   string `json:"uniqcode"`
	BornDay    int    `json:"born_day"`
	BornMonth  int    `json:"born_month"`
	BornYear   int    `json:"born_year"`
	PostalCode string `json:"postal_code"`
}
