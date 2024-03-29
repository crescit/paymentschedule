package types

import "time"

type Terms int64

const (
	net          Terms = 1
	installments       = 2
)

type PaymentInput struct {
	Amount        int
	FeePercentage int
	StartDate     time.Time
	Duration      int
	Terms         Terms
	Currency      string
}
