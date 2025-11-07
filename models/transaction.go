package models

import "time"

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	WalletID  uint    `gorm:"wallet_id"`
	Amount    float64 `gorm:"amount"`
	Type      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
