package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

// CreateAccount creating user account with balance = 0
func (s *Service) CreateAccount(userId uuid.UUID) error {
	err := s.db.Create(Account{
		Id:      uuid.New(),
		OwnerID: userId,
		Balance: 0,
	}).Error
	return err
}

// GetAccount return current balance for user
func (s *Service) GetAccount(userId uuid.UUID) (account *Account, err error) {
	err = s.db.Model(account).Where("owner_id = ?", userId).First(&account).Error
	return
}

// ChangeBalance changing balance for user
func (s *Service) ChangeBalance(account *Account) (err error) {
	err = s.db.Model(&account).Update("balance", account.Balance).Error
	return
}

type Account struct {
	Id      uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	OwnerID uuid.UUID `json:"ownerId" gorm:"type:uuid; not null; unique"`
	Balance int       `json:"balance"`
}
