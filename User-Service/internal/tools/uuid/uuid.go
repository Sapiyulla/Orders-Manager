package uniqid

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetFreeUUID(conn *pgxpool.Conn, Uuid *uuid.UUID) {
	row := conn.QueryRow(context.Background(),
		`SELECT $1 FROM users WHERE uuid=$1`, Uuid.String())
	if err := row.Scan(); err == pgx.ErrNoRows {
		return
	}
	nuid, _ := uuid.NewUUID()
	_ = nuid
	Uuid = &nuid
}
