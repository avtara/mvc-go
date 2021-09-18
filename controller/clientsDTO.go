package controller

import (
	"github.com/avtara/mvc-go/models"
	"gorm.io/gorm"
	"time"
)

type ClientDetailResponse struct {
	ID           uint64         `json:"client_id"`
	ClientKey    string         `json:"client_key"`
	ClientSecret string         `json:"client_secret"`
	Status       bool           `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" `
}

func ModelToClientDetailResponse(data *models.Client) *ClientDetailResponse {
	return &ClientDetailResponse{
		ID:           data.ID,
		ClientKey:    data.ClientKey,
		ClientSecret: data.ClientSecret,
		Status:       data.Status,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}
}

type PostClientRequest struct {
	ClientKey    string `json:"client_key" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	Status       bool   `json:"status"`
}

type PutClientRequest struct {
	ID           uint64 `json:"client_id" validate:"required"`
	ClientKey    string `json:"client_key" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	Status       bool   `json:"status"`
}

func PostClientsRequestToModelRequest(data *PostClientRequest) *models.Client {
	return &models.Client{
		ClientKey:    data.ClientKey,
		ClientSecret: data.ClientSecret,
		Status:       data.Status,
	}
}

func PutClientsRequestToModelRequest(data *PutClientRequest) *models.Client {
	return &models.Client{
		ID:           data.ID,
		ClientKey:    data.ClientKey,
		ClientSecret: data.ClientSecret,
		Status:       data.Status,
	}
}
