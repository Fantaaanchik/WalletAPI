package repository

import (
	"go.mod/config"
	"go.mod/models"
	"time"
)

type WalletRepository struct{}

func (r *WalletRepository) FindByID(id uint) (*models.Wallet, error) {
	var wallet models.Wallet
	err := config.DB.First(&wallet, id).Error
	return &wallet, err
}

func (r *WalletRepository) UpdateWallet(wallet *models.Wallet) error {
	return config.DB.Save(wallet).Error
}

func (r *WalletRepository) CreateTransaction(walletID uint, amount float64) error {
	trx := models.Transaction{
		WalletID:  walletID,
		Amount:    amount,
		Type:      "topup",
		CreatedAt: time.Now(),
	}

	return config.DB.Create(&trx).Error
}

func (r *WalletRepository) GetMonthlyTransactions(walletID uint, month int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := config.DB.Where("wallet_id = ? AND EXTRACT(MONTH FROM created_at) = ?", walletID, month).Find(&transactions).Error

	return transactions, err
}
