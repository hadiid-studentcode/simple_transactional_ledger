package models


type Entry struct {
	Id int64  `json:"id"`
	AccountId int64  `json:"account_id"`
	Amount float64 `json:"amount"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}