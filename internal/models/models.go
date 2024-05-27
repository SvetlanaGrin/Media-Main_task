package models

type Money struct{
	Amount int       `json:"amount"`
	Banknotes []int  `json:"banknotes"`
}

type Answer struct {
	Exchanges [][]int `json:"exchanges"`
}
		