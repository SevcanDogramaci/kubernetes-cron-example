package failed_transaction

import (
	"fmt"
)

type FailedTransactionService struct {
}

func (h FailedTransactionService) RetryTransaction() {
	fmt.Println("Retrying transaction...")
}
