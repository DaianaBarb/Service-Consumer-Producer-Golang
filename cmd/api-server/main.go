package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"project-golang/internal/fx/server"
)

// install swagger go install github.com/swaggo/swag/cmd/swag@latest
//export PATH=$PATH:$(go env GOPATH)/bin
//export source ~/.bashrc
// rodar swag init --output ./internal/docs --parseDependency na raiz do projeto se nao ele nao reconhce os handlrs

// @title DOC com o  Swagger
// @version 1.0
// @description Esta é uma API de simulação de emprestimos documentada com Swagger em Go.
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte
// @contact.email daiana.soares@hotmail.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1/simulation
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
