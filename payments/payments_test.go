package payments

import (
	"testing"

	inputs "github.com/crescit/paymentschedule/types"
)

func TestHandlePayment(t *testing.T) {
	tPaymentNil := inputs.PaymentInput{
		Terms: 1,
	}
	testPayments := make([]inputs.PaymentInput, 2)
	testPayments[0] = tPaymentNil
	for _, testPayment := range testPayments {
		err := HandlePayment(testPayment)
		if err != nil {
			t.Errorf("error encountered from handle payment")
		}
	}
}
