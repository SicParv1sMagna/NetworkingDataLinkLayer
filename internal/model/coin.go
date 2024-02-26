package model

type Coin struct {
	Name         string  `json:"name"`          // имя криптовалюты
	Symbol       string  `json:"symbol"`        // символ
	CurrentPrice float64 `json:"current_price"` // текущая стоимость
}
