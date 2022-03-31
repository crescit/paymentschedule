package payments

import (
	"testing"
	"time"

	inputs "github.com/crescit/paymentschedule/types"
	types "github.com/crescit/paymentschedule/types"
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
		Duration:      45,
		Currency:      "usd",
	}

	installmentPayment := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         2,
		StartDate:     date,
		Duration:      45,
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
		output, err := HandlePayment(testPayment)
		if err != nil && test1.wantErr[idx] != true {
			t.Errorf("error encountered from handle payment, err = %v, output = %v ", err, output)
		}
	}
}

func TestHandleNetPayment(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2022-01-12")
	if err != nil {
		t.Errorf("invalid date provided to testhandlepayment")
		return
	}
	date2, err := time.Parse("2006-01-02", "2022-01-10")
	if err != nil {
		t.Errorf("invalid date provided to testhandlepayment")
		return
	}
	netPayment45 := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         1,
		StartDate:     date,
		Duration:      45,
		Currency:      "usd",
	}
	output, err := HandleNetPayment(netPayment45)
	if err != nil {
		t.Errorf("error encountered during handlenetpayment err = %v", err)
		return
	}

	expected := types.DueOutput{
		Date:     "2022-02-28",
		Amount:   3150,
		Currency: "usd",
	}

	if len(output) < 1 || len(output) > 1 || expected != output[0] {
		t.Errorf("failed to get expected result for net45, got = %v, received = %v", output[0], expected)
		return
	}

	netPayment60 := inputs.PaymentInput{
		Amount:        3000,
		FeePercentage: 5,
		Terms:         1,
		StartDate:     date2,
		Duration:      60,
		Currency:      "usd",
	}
	output, err = HandleNetPayment(netPayment60)
	if err != nil {
		t.Errorf("error encountered during handlenetpayment err = %v", err)
		return
	}

	expected = types.DueOutput{
		Date:     "2022-03-11",
		Amount:   3150,
		Currency: "usd",
	}

	if len(output) < 1 || len(output) > 1 || expected != output[0] {
		t.Errorf("failed to get expected result for net45")
		return
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
		Duration:      60,
		Currency:      "usd",
	}
	output, err := HandleInstallmentPayment(installmentPayment)
	if err != nil {
		t.Errorf("error encountered during handleinstallmentpayment err = %v", err)
	}

	if len(output) != 3 || output[0].Amount != 1050 || output[0].Date != "2022-10-01" {
		t.Errorf("failed to get expected result for intall60, output = %v", output)
		return
	}
}
