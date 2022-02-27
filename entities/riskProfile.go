package entities

type RiskProfile struct {
	UserId       int     `json:"-"`
	BondPercent  float32 `json:"bond_percent"`
	StockPercent float32 `json:"stock_percent"`
	MMPercent    float32 `json:"mm_percent"`
}
