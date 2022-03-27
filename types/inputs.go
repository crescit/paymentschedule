package inputs

import "time"

type Terms int64

const (
	net          Terms = 0
	installments       = 1
)

type InstallmentInput struct {
	amount        int
	feePercentage int
	startDate     time.Time
	duration      time.Duration
	terms         Terms
	currency      string
}
