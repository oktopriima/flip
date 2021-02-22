/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21:46
**/

package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/flip/database/models"
)

type WithdrawalServices interface {
	Insert(withdrawals *models.Withdrawals, tx *gorm.DB) (*models.Withdrawals, error)
	FindOne(criteria map[string]interface{}) (*models.Withdrawals, error)
	Update(withdrawals *models.Withdrawals, tx *gorm.DB) error
}

type withdrawalServices struct {
	db *gorm.DB
}

func (w *withdrawalServices) Update(withdrawals *models.Withdrawals, tx *gorm.DB) error {
	if err := tx.Save(withdrawals).Error; err != nil {
		return err
	}
	return nil
}

func (w *withdrawalServices) FindOne(criteria map[string]interface{}) (*models.Withdrawals, error) {
	m := new(models.Withdrawals)
	res := w.db.Where(criteria).Find(&m).Scan(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (w *withdrawalServices) Insert(withdrawals *models.Withdrawals, tx *gorm.DB) (*models.Withdrawals, error) {
	db := tx.Create(&withdrawals)
	m := new(models.Withdrawals)
	if err := db.Error; err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(byteData, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func NewWithdrawalServices(db *gorm.DB) WithdrawalServices {
	return &withdrawalServices{db: db}
}
