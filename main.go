package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lilin1013/iban-validation/iban"
)

func main() {
	router := gin.Default()
	router.POST("/valid", iban.ValidIbanHandler)
	router.Run()
}
