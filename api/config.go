package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	account_sid := os.Getenv("TWILIO_ACCOUNT_SID")
	if account_sid == "" {
		log.Fatal("You must set your 'TWILIO_ACCOUNT_SID' environmental variable")
	}
	return account_sid
}

func envAUTHTOKEN() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	auth_token := os.Getenv("TWILIO_AUTHTOKEN")
	if auth_token == "" {
		log.Fatal("You must set your 'TWILIO_ACCOUNT_SID' environmental variable")
	}
	return auth_token
}

func envSERVICESID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	service_sid := os.Getenv("TWILIO_SERVICE_SID")
	if service_sid == "" {
		log.Fatal("You must set your 'TWILIO_ACCOUNT_SID' environmental variable")
	}
	return service_sid
}
