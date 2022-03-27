package inputs

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
	Duration      time.Duration
	Terms         Terms
	Currency      string
}
