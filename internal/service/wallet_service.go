package service

import (
	"errors"
	"go.mod/internal/repository"
	"time"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(r *repository.WalletRepository) *WalletService {
	return &WalletService{repo: r}
}

func (s *WalletService) CheckWalletExist(id uint) bool {
	_, err := s.repo.FindByID(id)
	return err == nil
}

func (s *WalletService) TopUp(id uint, amount float64) error {
	wallet, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("wallet not found")
	}

	limit := 100000.0
	if wallet.Type == "unidentified" {
		limit = 10000.0
	}

	if wallet.Balance+amount > limit {
		return errors.New("balance limit exceeded")
	}

	wallet.Balance += amount
	if err := s.repo.UpdateWallet(wallet); err != nil {
		return err
	}

	return s.repo.CreateTransaction(wallet.ID, amount)
}

func (s *WalletService) GetMonthlyStats(walletID uint) (int, float64, error) {
	transactions, err := s.repo.GetMonthlyTransactions(walletID, int(time.Now().Month()))
	if err != nil {
		return 0, 0, err
	}

	count := len(transactions)
	total := 0.0

	for _, t := range transactions {
		total += t.Amount
	}

	return count, total, err
}

func (s *WalletService) GetBalance(walletID uint) (float64, error) {
	wallet, err := s.repo.FindByID(walletID)
	if err != nil {
		return 0, errors.New("wallet not found")
	}

	return wallet.Balance, err
}
