package main

import "github.com/SevcanDogramaci/kubernetes-cron-example/internal/failed_transaction"

var failedTransactionHandler = failed_transaction.FailedTransactionService{}

func main() {
	failedTransactionHandler.RetryTransaction()
}
