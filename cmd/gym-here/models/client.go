package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Name       string   `gorm:"size:255;not null" json:"name"`
	Age        int      `gorm:"not null" json:"age"`
	TrainingID uint     `gorm:"not null" json:"training_id"`
	Training   Training `gorm:"foreignkey:TrainingID" json:"training"`
}

func GetAllClients(db *gorm.DB) ([]Client, error) {
	var clients []Client
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func GetClientByID(db *gorm.DB, id uint) (Client, error) {
	var client Client
	if err := db.First(&client, id).Error; err != nil {
		return Client{}, errors.New("Client not found!")
	}
	return client, nil
}

func CreateClient(db *gorm.DB, client *Client) error {
	if err := db.Create(client).Error; err != nil {
		return err
	}
	return nil
}

func UpdateClient(db *gorm.DB, client *Client) error {
	if err := db.Save(client).Error; err != nil {
		return err
	}
	return nil
}

func DeleteClient(db *gorm.DB, id uint) error {
	if err := db.Delete(&Client{}, id).Error; err != nil {
		return err
	}
	return nil
}
