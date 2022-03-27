package payments

import (
	"testing"
	"time"

	inputs "github.com/crescit/paymentschedule/types"
)

func TestHandlePayment(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2022-01-10")
	if err != nil {
		t.Errorf("invalid date provided to testhandlepayment")
	}
	var nilPayment inputs.PaymentInput

	netPayment := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         1,
		StartDate:     date,
		Duration:      time.Hour * 24 * 45,
		Currency:      "usd",
	}

	installmentPayment := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         2,
		StartDate:     date,
		Duration:      time.Hour * 24 * 45,
		Currency:      "usd",
	}

	testPayments := make([]inputs.PaymentInput, 3)
	testPayments[0] = netPayment
	testPayments[1] = installmentPayment
	testPayments[2] = nilPayment

	type tests struct {
		testPayments []inputs.PaymentInput
		wantErr      []bool
	}

	errorList := []bool{false, false, true}
	test1 := tests{
		testPayments: testPayments,
		wantErr:      errorList,
	}

	for idx, testPayment := range test1.testPayments {
		err := HandlePayment(testPayment)
		if err != nil && test1.wantErr[idx] != true {
			t.Errorf("error encountered from handle payment, err = %v", err)
		}
	}
}

func TestHandleNetPayment(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2022-01-10")
	if err != nil {
		t.Errorf("invalid date provided to testhandlepayment")
	}
	netPayment := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         1,
		StartDate:     date,
		Duration:      time.Hour * 24 * 45,
		Currency:      "usd",
	}
	_, err = HandleNetPayment(netPayment)
	if err != nil {
		t.Errorf("error encountered during handlenetpayment err = %v", err)
	}

}

func TestHandleInstallmentPayment(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2022-01-10")
	if err != nil {
		t.Errorf("invalid date provided to testhandlepayment")
	}
	installmentPayment := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         2,
		StartDate:     date,
		Duration:      time.Hour * 24 * 45,
		Currency:      "usd",
	}
	_, err = HandleNetPayment(installmentPayment)
	if err != nil {
		t.Errorf("error encountered during handleinstallmentpayment err = %v", err)
	}
}
