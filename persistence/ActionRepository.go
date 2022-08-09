package persistence

import (
	logger "github.com/SpirosOrfanos/readfromany/utils"
	"github.com/jmoiron/sqlx"
)

type ActionRepositoryDb struct {
	client *sqlx.DB
}

type ActionRepository interface {
	FindBy(req []string)
}

func (repo ActionRepositoryDb) FindBy(req []string) {
	sqlGetAccount := "SELECT action_id, type from actions where actionId = ?"
	var account Action
	err := repo.client.Get(&account, sqlGetAccount, req[0])
	if err != nil {
		logger.Info("Error while fetching account information: " + err.Error())

	}
}
