# paymentschedule
A function in my preferred programming language (Go) that generates the
estimated payment plan schedule for a set of parameters. Function may be added to an iFrame. 

# Running
To execute run go run main.go, this will open up a web service running on port 8080
Hit localhost:8080/payment with a POST request with the body following the following struct:
const (
	net          Terms = 1
	installments       = 2
)

type PaymentInput struct {
	Amount        int
	FeePercentage int
	StartDate     time.Time
	Duration      int
	Terms         Terms -> where terms is either 1 or 2 matching the above
	Currency      string
}



# Testing
To test the payments package run go test -v ./payments/ in the main directory

# Licensing
As the original author of this work, I Josue Jaquez hold full legal ownership of the work therein this repository. Any attempt to use it for commercial purposes is hereby restricted under Creative Commons Attribution-NonCommercial 2.0. The point of this work is merely demonstrative and educational and neither the author or any other entities have the right to use it for commercial purposes. 

https://creativecommons.org/licenses/by-nc/2.0/legalcode
