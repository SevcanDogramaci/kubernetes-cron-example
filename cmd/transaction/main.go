package main

import (
	"github.com/SevcanDogramaci/kubernetes-cron-example/internal/transaction"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	transactionHandler := transaction.TransactionHandler{}

	e.POST("/transaction", transactionHandler.CreateTransaction)
	e.Logger.Fatal(e.Start(":8080"))
}
