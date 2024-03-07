package client

type Client struct {
	ID      int
	Limit   int
	Balance int
}

type TransactionBalance struct {
	Value       int
	Type        string
	Description string
	Date        string
}

type BalanceWithTransactions struct {
	Limit        int
	Balance      int
	Transactions []TransactionBalance
}

type ClientRepositoryInterface interface {
	Exists(id int) (bool, error)
	IncreaseBalance(id int, value int) (*Client, error)
	DecreaseBalance(id int, value int) (*Client, error)
	BalanceWithTransactions(id int, limitTransactions int) (*BalanceWithTransactions, error)
}
