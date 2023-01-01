package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/satori/go.uuid"
)

func (u Account) TableName() string {
	return "accounts"
}

type Account struct {
	Id      uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	OwnerID uuid.UUID `json:"ownerId" gorm:"type:uuid; not null; unique"`
	Balance int       `json:"balance"`
}

func init() {

	mx := &gormigrate.Migration{
		ID:       "0002",
		Migrate:  mixture.CreateTableM(&Account{}),
		Rollback: mixture.DropTableR(&Account{}),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
