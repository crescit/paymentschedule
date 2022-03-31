package payments

import (
	"encoding/json"
	"fmt"

	types "github.com/crescit/paymentschedule/types"
)

// HandlePayment either returns an error if it receives a valid term or calls the terms respective helper func
func HandlePayment(payment types.PaymentInput) (string, error) {
	if payment.Terms < 1 || payment.Terms >= 3 {
		return "", fmt.Errorf("invalid terms received from payment struct")
	}
	if payment.Terms == 1 {
		output, err := HandleNetPayment(payment)
		if err != nil {
			return "", err
		}
		return PrintInJson(output), nil
	}
	output, err := HandleInstallmentPayment(payment)
	if err != nil {
		return "", err
	}
	return PrintInJson(output), nil
}

func HandleNetPayment(payment types.PaymentInput) ([]types.DueOutput, error) {
	var netPayments []types.DueOutput
	startVal := payment.Amount
	endDate := payment.StartDate.AddDate(0, 0, payment.Duration)
	interest := float64(payment.Amount) * (float64(payment.FeePercentage) / 100)
	var netPayment types.DueOutput
	netPayment.Amount = startVal + int(interest)
	netPayment.Currency = payment.Currency
	netPayment.Date = endDate.Format("2006-02-01")
	netPayments = append(netPayments, netPayment)
	return netPayments, nil
}

func HandleInstallmentPayment(payment types.PaymentInput) ([]types.DueOutput, error) {
	totalPayments := 1 + (payment.Duration / 30 % 30)
	output := make([]types.DueOutput, totalPayments)
	interest := float64(payment.Amount) * (float64(payment.FeePercentage) / 100)
	total := int(payment.Amount/totalPayments) + int(interest)/totalPayments
	newDate := payment.StartDate
	for i := 0; i < totalPayments; i++ {
		var installmentPayment types.DueOutput
		installmentPayment.Amount = total
		installmentPayment.Currency = payment.Currency
		installmentPayment.Date = newDate.Format("2006-02-01")
		output[i] = installmentPayment
		newDate = newDate.AddDate(0, 0, 30)
	}
	return output, nil
}

// write up is confusing 45 days after 1-12 would be 2-26 adjusting for business days doesn't give 2-28 either???
// func AddBusinessDays(startDate time.Time, durationInDays int) time.Time {
// 	t := startDate
// 	days := 0
// 	for {
// 		fmt.Printf("%v  _ %v \n", days, durationInDays)
// 		if days >= durationInDays {
// 			break
// 		}
// 		if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
// 			days++
// 		}
// 		t = t.Add(time.Hour * 24)
// 	}
// 	return t
// }

func PrintInJson(dueoutput []types.DueOutput) string {
	output, err := json.Marshal(dueoutput)
	if err != nil {
		fmt.Println(err)
	}
	return string(output)
}
