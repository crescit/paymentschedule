package payments

import (
	"fmt"

	types "github.com/crescit/paymentschedule/types"
)

func HandlePayment(payment types.PaymentInput) error {
	if payment.Terms < 1 || payment.Terms >= 3 {
		return fmt.Errorf("invalid terms received from payment struct")
	}
	return nil
}

func HandleNetPayment(payment types.PaymentInput) ([]types.DueOutput, error) {
	return nil, nil
}

func HandleInstallmentPayment(payment types.PaymentInput) ([]types.DueOutput, error) {
	return nil, nil
}
