package cmd

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rafaqwe1/rinha-backend-2024/domain/client"
	"github.com/rafaqwe1/rinha-backend-2024/domain/shared"
	"github.com/rafaqwe1/rinha-backend-2024/domain/transaction"
	check_balance "github.com/rafaqwe1/rinha-backend-2024/usecases/check-balance"
	create_transaction "github.com/rafaqwe1/rinha-backend-2024/usecases/create-transaction"
)

type App struct {
	CheckBalanceUseCase      *check_balance.CheckBalanceUseCase
	CreateTransactionUseCase *create_transaction.CreateTransactionUseCase
}

// Simulation RinhaBackendCrebitosSimulation completed in 245 seconds
func Execute(clientRepository client.ClientRepositoryInterface, transactionRepository transaction.TransactionRepositoryInterface) {

	apiFiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	checkBalanceUseCase := check_balance.NewCheckBalanceUseCase(clientRepository, transactionRepository)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(clientRepository, transactionRepository)

	app := App{CheckBalanceUseCase: checkBalanceUseCase, CreateTransactionUseCase: createTransactionUseCase}

	apiFiber.Post("/clientes/:id/transacoes", app.PostTransaction)
	apiFiber.Get("/clientes/:id/extrato", app.CheckBalance)

	apiFiber.Listen(":8080")
}

func (app *App) CheckBalance(c *fiber.Ctx) error {

	clientIdInt, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("O parâmetro 'id' deve ser um número inteiro")
	}

	output, err := app.CheckBalanceUseCase.Execute(check_balance.Input{ClientId: clientIdInt})

	if err != nil {
		httpCode := app.getHttpCodeByError(err)
		if httpCode == fiber.StatusInternalServerError {
			return c.SendStatus(httpCode)
		}
		return c.Status(httpCode).JSON(&fiber.Map{"error": err.Error()})
	}

	return c.JSON(output)
}

func (app *App) PostTransaction(c *fiber.Ctx) error {

	clientIdInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity).SendString("O parâmetro 'id' deve ser um número inteiro")
	}

	var input create_transaction.Input
	err = c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Erro ao decodificar o JSON")
	}

	input.ClientId = clientIdInt
	output, err := app.CreateTransactionUseCase.Execute(input)

	if err != nil {
		httpCode := app.getHttpCodeByError(err)
		if httpCode == http.StatusInternalServerError {
			return c.SendStatus(httpCode)
		}

		return c.Status(httpCode).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(output)
}

func (app *App) getHttpCodeByError(err error) int {

	if errors.Is(err, shared.NotFoundError) {
		return fiber.StatusNotFound
	}

	if errors.As(err, &shared.TypeValidationError) {
		return fiber.StatusUnprocessableEntity
	}

	log.Println(err.Error())
	return fiber.StatusInternalServerError
}
