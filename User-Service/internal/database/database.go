package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote: true,
	})
}

func NewPool(DSN string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		logrus.Panicln(err.Error())
		return nil
	}
	return pool
}

func NewConnect(pool *pgxpool.Pool) *pgxpool.Conn {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		logrus.Warnln(err.Error())
		return nil
	}
	return conn
}
