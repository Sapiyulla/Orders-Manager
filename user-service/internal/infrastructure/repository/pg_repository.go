package repository

import (
	"context"
	"time"
	"user-service/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type PGAccountRepository struct {
	pool *pgxpool.Pool
}

func NewPGUserRepository(pool *pgxpool.Pool) *PGAccountRepository {
	return &PGAccountRepository{pool: pool}
}

func (pg *PGAccountRepository) Add(account *domain.Account) error {
	conn, err := pg.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	hashPS, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account.Password = string(hashPS)
	_, err = conn.Exec(context.Background(),
		`INSERT INTO users VALUES ($1, $2, $3, $4)`,
		account.UUID, account.Login, account.Password, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	return nil
}

func (pg *PGAccountRepository) Presence(account *domain.Account) (bool, error) {
	conn, err := pg.pool.Acquire(context.Background())
	if err != nil {
		return false, err
	}
	defer conn.Release()
	_field := ""
	row := conn.QueryRow(context.Background(),
		`SELECT 1 FROM users WHERE login=$1`, account.Login)
	if err := row.Scan(&_field); err != nil && err != pgx.ErrNoRows {
		return false, err
	} else if err != nil && err == pgx.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func (pg *PGAccountRepository) Get(login, password string) (*domain.Account, error) {
	conn, err := pg.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(),
		`SELECT password FROM users WHERE login=$1`, login)
	row_password := ""
	if err := row.Scan(&row_password); err != nil {
		return nil, err
	}
	var account domain.Account
	if err := bcrypt.CompareHashAndPassword([]byte(row_password), []byte(password)); err != nil {
		return nil, domain.ErrInvalidPassword
	} else {
		newRow := conn.QueryRow(context.Background(),
			`SELECT uuid, login FROM users WHERE login=$1`, login)
		newRow.Scan(&account.UUID, &account.Login)
		return &account, nil
	}
}
