package transaction

import (
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
)

type Transaction struct {
	ID          int
	ClientId    int
	Description string
	Type        string
	Value       int
	Date        string
}

const (
	Credit = "c"
	Debit  = "d"
)

func (t *Transaction) Validate() error {

	if len(t.Description) == 0 {
		return shared.NewValidationError("Description is required")
	}

	if len(t.Description) > 10 {
		return shared.NewValidationError("Description maximum size is 10")
	}

	if t.Type != Credit && t.Type != Debit {
		return shared.NewValidationError("Invalid type")
	}

	if t.Value <= 0 {
		return shared.NewValidationError("The value must be greater than 0")
	}

	return nil
}

type TransactionRepositoryInterface interface {
	Add(entity Transaction) error
}
