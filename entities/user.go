package entities

type User struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Age         int          `json:"age"`
	RiskProfile *RiskProfile `gorm:"foreignKey:UserId" json:"risk_profile,omitempty"`
}
