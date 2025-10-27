
package models

import "time"

type Account struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Status    string    `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	AccountID  string    `gorm:"index;not null" json:"account_id"`
	Symbol     string    `gorm:"index" json:"symbol"`
	Side       string    `json:"side"`
	Qty        int       `json:"qty"`
	Price      float64   `json:"price"`
	Status     string    `gorm:"index" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type LedgerEntry struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	AccountID  string    `gorm:"index;not null" json:"account_id"`
	Type       string    `json:"type"`
	Amount     float64   `json:"amount"`
	ImmutHash  string    `json:"immuthash"`
	CreatedAt  time.Time `json:"created_at"`
}

type AuditLog struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ActorID   *string   `json:"actor_id"`
	Action    string    `gorm:"index" json:"action"`
	IP        string    `json:"ip"`
	Device    string    `json:"device"`
	Success   bool      `json:"success"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}
