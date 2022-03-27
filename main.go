package main

import (
	payments "github.com/crescit/paymentschedule/payments"
	types "github.com/crescit/paymentschedule/types"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var nilPayment types.PaymentInput
	payments.HandlePayment(nilPayment)
}
