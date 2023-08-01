package dto

import (
	"github.com/google/uuid"
)

type DeleteFromCartRequest struct {
	UserId uuid.UUID `json:"userId"`
	ItemId string    `json:"itemId"`
}

type DeleteFromCartResponse struct {
	Err    error     `json:"error,omitempty"`
	UserId uuid.UUID `json:"userId,omitempty"`
	ItemId string    `json:"itemId,omitempty"`
}
