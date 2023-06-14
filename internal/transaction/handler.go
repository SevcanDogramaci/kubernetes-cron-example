package transaction

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
}

func (h TransactionHandler) CreateTransaction(ctx echo.Context) error {
	fmt.Println("Creating transaction...")
	return nil
}
