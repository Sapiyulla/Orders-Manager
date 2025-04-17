package database

import (
	"context"
	"user-service/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(conf *config.Database, password string) (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://"+conf.Admin+":"+password+"@"+conf.Host+":"+conf.Port+"/"+conf.DB_name)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	p.Pool.Close()
}
