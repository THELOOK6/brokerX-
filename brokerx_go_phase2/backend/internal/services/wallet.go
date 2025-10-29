package services

import (
	"brokerx/backend/internal/models"
	"brokerx/backend/internal/store"
	"errors"
)

func Deposit(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	entry := models.LedgerEntry{
		AccountID: accountID,
		Type:      "DEPOSIT",
		Amount:    amount,
	}

	return store.DB.Create(&entry).Error
}

func GetBalance(accountID string) (float64, error) {
	var balance float64
	err := store.DB.Model(&models.LedgerEntry{}).
		Where("account_id = ?", accountID).
		Select("COALESCE(SUM(amount), 0)").Scan(&balance).Error

	return balance, err
}
