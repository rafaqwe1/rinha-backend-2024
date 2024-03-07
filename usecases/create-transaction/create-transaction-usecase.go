package create_transaction

import (
	"errors"

	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
	"github.com/rafaqwe1/rinha-backend-2024/domain/transaction"
)

type CreateTransactionUseCase struct {
	ClientRepository      client.ClientRepositoryInterface
	TransactionRepository transaction.TransactionRepositoryInterface
}

func NewCreateTransactionUseCase(clientRepository client.ClientRepositoryInterface, transactionRepository transaction.TransactionRepositoryInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		ClientRepository:      clientRepository,
		TransactionRepository: transactionRepository,
	}
}

type Output struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type Input struct {
	ClientId    int
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (useCase *CreateTransactionUseCase) Execute(input Input) (*Output, error) {

	clientExists, err := useCase.ClientRepository.Exists(input.ClientId)
	if err != nil {
		return &Output{}, err
	}

	if !clientExists {
		return &Output{}, shared.NotFoundError
	}

	entity := transaction.Transaction{
		ClientId:    input.ClientId,
		Description: input.Description,
		Type:        input.Type,
		Value:       input.Value,
	}

	if err = entity.Validate(); err != nil {
		return &Output{}, err
	}

	var clientEntity *client.Client
	if entity.Type == transaction.Debit {
		clientEntity, err = useCase.ClientRepository.DecreaseBalance(entity.ClientId, entity.Value)
	} else {
		clientEntity, err = useCase.ClientRepository.IncreaseBalance(entity.ClientId, entity.Value)
	}

	if err != nil && errors.Is(err, shared.NoRowsAffectedError) {
		return &Output{}, shared.NewValidationError("Insufficient limit")
	}

	if err != nil {
		return &Output{}, err
	}

	err = useCase.TransactionRepository.Add(entity)
	if err != nil {
		return &Output{}, err
	}

	return &Output{
		Limit:   clientEntity.Limit,
		Balance: clientEntity.Balance,
	}, nil
}
