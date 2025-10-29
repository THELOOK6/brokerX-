package services

import (
	"brokerx/backend/internal/models"
	"brokerx/backend/internal/store"
	"errors"
)

func PlaceOrder(accountID, symbol, side string, qty int, price float64) (models.Order, error) {
	if qty <= 0 || price <= 0 {
		return models.Order{}, errors.New("invalid qty or price")
	}

	// Calculate balance
	balance, err := GetBalance(accountID)
	if err != nil {
		return models.Order{}, err
	}

	cost := float64(qty) * price

	if side == "BUY" && balance < cost {
		return models.Order{}, errors.New("insufficient funds")
	}

	// Deduct funds
	if side == "BUY" {
		store.DB.Create(&models.LedgerEntry{
			AccountID: accountID,
			Type:      "TRADE",
			Amount:    -cost,
		})
	}

	order := models.Order{
		AccountID: accountID,
		Symbol:    symbol,
		Side:      side,
		Qty:       qty,
		Price:     price,
		Status:    "FILLED",
	}

	err = store.DB.Create(&order).Error
	return order, err
}

func GetOrders(accountID string) ([]models.Order, error) {
	var orders []models.Order
	err := store.DB.Where("account_id = ?", accountID).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}
