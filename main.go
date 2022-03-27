package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/crescit/paymentschedule/payments"
	"github.com/crescit/paymentschedule/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.POST("/payment", HandlePaymentEndpoint)
	r.Run()
}

func HandlePaymentEndpoint(c *gin.Context) {
	var Payment types.PaymentInput
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Print("Unable to bind request body")
		c.AbortWithStatus(400)
		return
	}
	json.Unmarshal(body, &Payment)
	response, err := payments.HandlePayment(Payment)
	if err != nil {
		log.Printf("unable to process payment, %v", err)
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusOK, response)
}
