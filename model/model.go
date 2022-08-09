package model

type Account struct {
	AccountId string `db:"action_id"`
	Type      string `db:"type"`
}
