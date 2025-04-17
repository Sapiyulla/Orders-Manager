package domain

import "errors"

type AccountRepository interface {
	Add(*Account) error
	Presence(*Account) (bool, error)
	Get(string, string) (*Account, error)
}

type Account struct {
	UUID     string `json:"uuid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var (
	ErrInvalidPassword = errors.New("invalid password")
)
