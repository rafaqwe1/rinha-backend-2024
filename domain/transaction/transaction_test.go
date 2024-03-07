package transaction_test

import (
	"testing"

	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
	"github.com/rafaqwe1/rinha-backend-2024/domain/transaction"
	"github.com/stretchr/testify/require"
)

func TestTransactionValidateValid(t *testing.T) {
	transaction := transaction.Transaction{
		ID:          1,
		ClientId:    1,
		Description: "test",
		Type:        "c",
		Value:       100,
	}

	err := transaction.Validate()
	require.Nil(t, err)
}

func TestTransactionInvalidValue(t *testing.T) {
	transaction := transaction.Transaction{
		ID:          1,
		ClientId:    1,
		Description: "test",
		Type:        "c",
		Value:       0,
	}

	err := transaction.Validate()
	require.Equal(t, "The value must be greater than 0", err.Error())
	require.ErrorAs(t, err, &shared.TypeValidationError)
}

func TestTransactionInvalidDescription(t *testing.T) {
	transaction := transaction.Transaction{
		ID:          1,
		ClientId:    1,
		Description: "test_more_than_ten_characteres",
		Type:        "c",
		Value:       1,
	}

	err := transaction.Validate()
	require.Equal(t, "Description maximum size is 10", err.Error())

	transaction.Description = ""
	err = transaction.Validate()
	require.Equal(t, "Description is required", err.Error())
	require.ErrorAs(t, err, &shared.TypeValidationError)
}

func TestTransactionInvalidType(t *testing.T) {
	transaction := transaction.Transaction{
		ID:          1,
		ClientId:    1,
		Description: "test",
		Type:        "a",
		Value:       1,
	}

	err := transaction.Validate()
	require.Equal(t, "Invalid type", err.Error())
	require.ErrorAs(t, err, &shared.TypeValidationError)
}
