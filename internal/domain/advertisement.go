package domain

import (
	"errors"
	"time"
)

type Advertisement struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Price       float32   `json:"price" binding:"required"`
	CreatedTime time.Time `json:"createdTime" db:"createdTime"`
	IsActive    bool      `json:"IsActive" db:"active"`
}

type UpdateAdvertisementInput struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	IsActive    *bool    `json:"isActive"`
}

func (input UpdateAdvertisementInput) ValidateUpdateInput() error {
	if input.Title == nil || *input.Title == "" {
		return errors.New("title is required")
	}
	if input.Price == nil || *input.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil

}
