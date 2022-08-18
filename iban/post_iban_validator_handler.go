package iban

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	IsValid bool   `json:"isValid"`
	Reason  string `json:"reason,omitempty" bson:",omitempty"`
}

type ibanRequestBody struct {
	IbanNumber string `json:"ibanNumber"`
}

func ValidIbanHandler(context *gin.Context) {

	var requestBody ibanRequestBody

	if err := context.BindJSON(&requestBody); err != nil {
		return
	}

	isValid, reason := isValidIbanNum(requestBody.IbanNumber)

	context.IndentedJSON(http.StatusOK, response{
		IsValid: isValid,
		Reason:  reason,
	})
}
