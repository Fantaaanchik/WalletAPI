package models

type Wallet struct {
	ID       uint `gorm:"primaryKey"`
	UserId   uint
	Type     string  `gorm:"type"` // "identified" | "unidentified"
	Balance  float64 `gorm:"got null; default:0"`
	Currency string  `gorm:"default:'TJS'"`
}
