package models

type Account struct {
	Id int64  `json:"id"`
	Name string `json:"name"`
	Balance float64 `json:"balance"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

