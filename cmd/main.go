package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peb-peb/golang-otp-verify-twilio/api"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	app.SMSRoutes()

	router.Run(":8000")
}
