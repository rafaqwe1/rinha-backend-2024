package check_balance_test

import (
	"testing"

	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	mock_client "github.com/rafaqwe1/rinha-backend-2024/domain/client/mocks"
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
	mock_transaction "github.com/rafaqwe1/rinha-backend-2024/domain/transaction/mock"
	check_balance "github.com/rafaqwe1/rinha-backend-2024/usecases/check-balance"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCheckBalanceNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)
	clientRep.EXPECT().BalanceWithTransactions(1, 10).Return(&client.BalanceWithTransactions{}, shared.NotFoundError)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)

	usecase := check_balance.NewCheckBalanceUseCase(clientRep, transactionRep)

	_, err := usecase.Execute(check_balance.Input{
		ClientId: 1,
	})

	require.ErrorIs(t, err, shared.NotFoundError)
}

func TestCheckBalanceFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)

	var transactions []client.TransactionBalance
	transactions = append(transactions, client.TransactionBalance{Description: "first", Type: "c", Value: 10})
	transactions = append(transactions, client.TransactionBalance{Description: "second", Type: "d", Value: 12})

	response := &client.BalanceWithTransactions{
		Limit:        10,
		Balance:      10,
		Transactions: transactions,
	}

	clientId := 1
	clientRep.EXPECT().BalanceWithTransactions(clientId, 10).Return(response, nil)

	usecase := check_balance.NewCheckBalanceUseCase(clientRep, transactionRep)

	output, err := usecase.Execute(check_balance.Input{ClientId: clientId})

	require.Nil(t, err)
	require.Equal(t, response.Limit, output.Balance.Limit)
	require.Equal(t, response.Balance, output.Balance.Total)
	require.Equal(t, len(transactions), len(output.LastTransactions))
}
