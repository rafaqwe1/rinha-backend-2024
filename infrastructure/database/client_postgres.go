package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
)

type ClientDbPostgres struct {
	db *pgxpool.Pool
}

func NewClientDbPostgres(db *pgxpool.Pool) *ClientDbPostgres {
	return &ClientDbPostgres{db: db}
}

func (c *ClientDbPostgres) Exists(id int) (bool, error) {
	var exists bool
	err := c.db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM clients WHERE id=$1)", id).Scan(&exists)
	return exists, err
}

func (c *ClientDbPostgres) DecreaseBalance(id int, value int) (*client.Client, error) {
	var limit, balance int
	err := c.db.QueryRow(context.Background(), "update clients set balance = balance - $1 where id = $2 and ABS(balance - $3) <= clients.limit RETURNING clients.limit, balance",
		value,
		id,
		value).Scan(&limit, &balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return &client.Client{}, shared.NoRowsAffectedError
	}

	if err != nil {
		return &client.Client{}, err
	}

	return &client.Client{ID: id, Limit: limit, Balance: balance}, nil
}

func (c *ClientDbPostgres) IncreaseBalance(id int, value int) (*client.Client, error) {
	var limit, balance int
	err := c.db.QueryRow(context.Background(), "update clients set balance = balance + $1 where id = $2 RETURNING clients.limit, balance",
		value,
		id).Scan(&limit, &balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return &client.Client{}, shared.NotFoundError
	}

	if err != nil {
		return &client.Client{}, err
	}

	return &client.Client{ID: id, Limit: limit, Balance: balance}, nil
}

func (c *ClientDbPostgres) BalanceWithTransactions(id int, limitTransactions int) (*client.BalanceWithTransactions, error) {
	rows, err := c.db.Query(context.Background(), `select "balance", "limit", "value", "type", description, date_added from clients left join transactions on clients.id = client_id where clients.id = $1 order by date_added desc limit $2`, id, limitTransactions)

	if errors.Is(err, pgx.ErrNoRows) {
		return &client.BalanceWithTransactions{}, shared.NotFoundError
	}

	if err != nil {
		return &client.BalanceWithTransactions{}, err
	}

	defer rows.Close()
	output := &client.BalanceWithTransactions{}
	for rows.Next() {
		var balance, limit, value sql.NullInt64
		var transactionType, description, date sql.NullString

		rows.Scan(&balance, &limit, &value, &transactionType, &description, &date)

		if !limit.Valid {
			return &client.BalanceWithTransactions{}, shared.NotFoundError
		}

		output.Balance = int(balance.Int64)
		output.Limit = int(limit.Int64)

		if transactionType.Valid {
			output.Transactions = append(output.Transactions, client.TransactionBalance{
				Value:       int(value.Int64),
				Type:        transactionType.String,
				Description: description.String,
				Date:        date.String,
			})
		}
	}

	if len(rows.RawValues()) == 0 {
		return &client.BalanceWithTransactions{}, shared.NotFoundError
	}

	return output, nil
}
