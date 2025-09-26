package model

import "time"

type Verification struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Token     string    `json:"token" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Used      bool      `json:"used" gorm:"default:false"`
}
