package model

type Action struct {
	ActionId string `db:"action_id"`
	Type     string `db:"type"`
}
