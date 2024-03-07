package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaqwe1/rinha-backend-2024/domain/transaction"
)

type TransactionPostgresRepository struct {
	db *pgxpool.Pool
}

func NewTransactionDbPostgres(db *pgxpool.Pool) *TransactionPostgresRepository {
	return &TransactionPostgresRepository{db: db}
}

func (repository *TransactionPostgresRepository) Add(entity transaction.Transaction) error {
	_, err := repository.db.Exec(context.Background(), "insert into transactions (client_id, type, description, value) values ($1, $2, $3, $4)",
		entity.ClientId, entity.Type, entity.Description, entity.Value)

	return err
}
