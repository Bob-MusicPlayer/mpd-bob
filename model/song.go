package model

type Song struct {
	Artist   string  `json:"Artist"`
	Duration float64 `json:"Duration"`
	Title    string  `json:"Title"`
}
