package database

import (
	"errors"
	"github.com/avtara/mvc-go/config"
	"github.com/avtara/mvc-go/models"
	"gorm.io/gorm"
)

func GetClients() (interface{}, error) {
	var client []models.Client
	err := config.DB.Find(&client).Error
	if err != nil {
		return nil, err
	}

	if len(client) <= 0 {
		return nil, errors.New("not found")
	}
	return client, nil
}

func GetClientByID(ID int) (*models.Client, error) {
	var client *models.Client
	err := config.DB.Where("id = ?", ID).First(&client).Error
	errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return nil, errors.New("not found")
	}

	return client, nil
}

func InsertClient(client *models.Client) (*models.Client, error) {
	err := config.DB.Create(&client)
	if err.Error != nil {
		return nil, err.Error
	}
	config.DB.Where("id = ?", client.ID).First(&client)

	return client, nil
}

func UpdateClient(client *models.Client) (*models.Client, error) {
	err := config.DB.Table("clients").Updates(client).Error

	if err != nil {
		return nil, errors.New("not found")
	}
	return client, nil
}

func DeleteClient(ID int) (*models.Client, error) {
	var client *models.Client
	err := config.DB.Delete(&client, ID).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return client, nil
}
