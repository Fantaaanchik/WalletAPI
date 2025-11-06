package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"name"`
	SecretKey string `gorm:"secretKey"`
	IsActive  bool
}
