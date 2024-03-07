package check_balance

import (
	"time"

	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	"github.com/rafaqwe1/rinha-backend-2024/domain/transaction"
)

type CheckBalanceUseCase struct {
	ClientRepository      client.ClientRepositoryInterface
	TransactionRepository transaction.TransactionRepositoryInterface
}

func NewCheckBalanceUseCase(clientRepository client.ClientRepositoryInterface, transactionRepository transaction.TransactionRepositoryInterface) *CheckBalanceUseCase {
	return &CheckBalanceUseCase{
		ClientRepository:      clientRepository,
		TransactionRepository: transactionRepository,
	}
}

type Input struct {
	ClientId int
}

type OutputTransactions struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	Date        string `json:"realizada_em"`
}

type Balance struct {
	Total int    `json:"total"`
	Date  string `json:"data_extrato"`
	Limit int    `json:"limite"`
}

type Output struct {
	Balance          Balance              `json:"saldo"`
	LastTransactions []OutputTransactions `json:"ultimas_transacoes"`
}

func (useCase *CheckBalanceUseCase) Execute(input Input) (Output, error) {
	balance, err := useCase.ClientRepository.BalanceWithTransactions(input.ClientId, 10)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		Balance: Balance{
			Total: balance.Balance,
			Limit: balance.Limit,
			Date:  time.Now().Format(time.RFC3339),
		},
	}

	for _, transaction := range balance.Transactions {
		output.LastTransactions = append(output.LastTransactions, OutputTransactions{
			Value:       transaction.Value,
			Type:        transaction.Type,
			Description: transaction.Description,
			Date:        transaction.Date,
		})
	}

	return output, nil
}
