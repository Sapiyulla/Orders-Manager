package application

import (
	"errors"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

type AccountUseCase struct {
	repo domain.AccountRepository
}

func NewAccountRepository(repo domain.AccountRepository) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

func (uc *AccountUseCase) Register(Uuid *string, login, password string) error {
	if *Uuid == "" {
		*Uuid = uuid.New().String()
	}
	account := &domain.Account{
		UUID:     *Uuid,
		Login:    login,
		Password: password,
	}
	if presence, err := uc.repo.Presence(account); err != nil {
		return err
	} else if presence {
		return errors.New("account with this data is already exists")
	}

	return uc.repo.Add(account)
}

func (uc *AccountUseCase) Login(login, password string) (*domain.Account, error) {
	if presence, err := uc.repo.Presence(&domain.Account{Login: login, Password: password}); err != nil && !errors.Is(err, domain.ErrInvalidPassword) {
		return nil, err
	} else if err == domain.ErrInvalidPassword {
		return nil, domain.ErrInvalidPassword
	} else if !presence {
		return nil, errors.New("account not found")
	}
	return uc.repo.Get(login, password)
}
