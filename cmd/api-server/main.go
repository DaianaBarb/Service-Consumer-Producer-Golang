package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"project-golang/internal/fx/server"
)

var (
	QUEUE_URL string
)

func LoadEnv() {
	QUEUE_URL = os.Getenv("QUEUE_URL")

}

func init() {
	if err := godotenv.Load("../../configs/prod.env"); err != nil && !os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"application": "consumer",
		}).Fatalf("Error loading .env - %s", err.Error())
	}
	LoadEnv()
}

func main() {
	fmt.Println("iniciando aplicação")
	fmt.Println("server.Start")


	server.Start()


}
