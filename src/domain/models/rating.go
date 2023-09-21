package models

import "time"

type Rating struct {
	ID string `json:"id" gorn:"primaryKey"`
	OrderID string `json:"order_id" gorm:"index:idx_order_rating"`
	Rating int `json:"rating"`
	Comment string `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	}