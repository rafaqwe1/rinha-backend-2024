package create_transaction_test

import (
	"testing"

	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	mock_client "github.com/rafaqwe1/rinha-backend-2024/domain/client/mocks"
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
	mock_transaction "github.com/rafaqwe1/rinha-backend-2024/domain/transaction/mock"
	create_transaction "github.com/rafaqwe1/rinha-backend-2024/usecases/create-transaction"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateClientNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)
	clientRep.EXPECT().Exists(1).Return(false, nil)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)

	usecase := create_transaction.NewCreateTransactionUseCase(clientRep, transactionRep)
	_, err := usecase.Execute(create_transaction.Input{ClientId: 1})

	require.ErrorIs(t, err, shared.NotFoundError)
}

func TestCreateInsuficientLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := &client.Client{
		ID:      1,
		Limit:   10,
		Balance: 10,
	}

	input := create_transaction.Input{
		ClientId:    1,
		Value:       10,
		Type:        "d",
		Description: "test",
	}

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)
	clientRep.EXPECT().Exists(1).Return(true, nil)
	clientRep.EXPECT().DecreaseBalance(client.ID, input.Value).Return(client, shared.NoRowsAffectedError)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)

	usecase := create_transaction.NewCreateTransactionUseCase(clientRep, transactionRep)
	_, err := usecase.Execute(input)

	require.ErrorAs(t, err, &shared.TypeValidationError)
}

func TestCreateInvalidData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := create_transaction.Input{
		ClientId: 1,
		Type:     "f",
	}

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)
	clientRep.EXPECT().Exists(input.ClientId).Return(true, nil)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)

	usecase := create_transaction.NewCreateTransactionUseCase(clientRep, transactionRep)
	_, err := usecase.Execute(input)

	require.ErrorAs(t, err, &shared.TypeValidationError)
}

func TestCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := &client.Client{
		ID:      1,
		Limit:   10,
		Balance: 10,
	}

	input := create_transaction.Input{
		ClientId:    1,
		Value:       10,
		Type:        "c",
		Description: "test",
	}

	clientRep := mock_client.NewMockClientRepositoryInterface(ctrl)
	clientRep.EXPECT().Exists(input.ClientId).Return(true, nil)
	clientRep.EXPECT().IncreaseBalance(client.ID, input.Value).Return(client, nil)

	transactionRep := mock_transaction.NewMockTransactionRepositoryInterface(ctrl)
	transactionRep.EXPECT().Add(gomock.Any()).Return(nil)

	usecase := create_transaction.NewCreateTransactionUseCase(clientRep, transactionRep)
	output, err := usecase.Execute(input)

	require.Nil(t, err)
	require.Equal(t, output.Balance, client.Balance)
	require.Equal(t, client.Limit, output.Limit)
}
