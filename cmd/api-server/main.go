package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"project-golang/internal/fx/server"
	//"github.com/spf13/viper"
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
			"application": "migration-worker",
			"brand":       "b2w",
		}).Fatalf("Error loading .env - %s", err.Error())
	}
	LoadEnv()

	//awssqs.Configure(os.Getenv("QUEUE_VPC_ENDPOINTS"))

	//cassandra.OpenDatabases()
}

func main() {
	fmt.Println("iniciando aplicação")
	fmt.Println("server.Start")
	var2 := os.Getenv("QUEUE_URL")
	fmt.Println(" estou imprimindo: "+var2)

	// viper.AddConfigPath("../../configs/worker")
	// viper.SetConfigName("development") // Nome do arquivo sem extensão
	// viper.SetConfigType("yaml")

	// viper.AddConfigPath("../../configs")
	// viper.SetConfigName("default") // Nome do arquivo sem extensão
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("../../configs")
	// viper.SetConfigName("rest") // Nome do arquivo sem extensão
	// viper.SetConfigType("yaml")

	server.Start()

	fmt.Println("terminando aplicação")

}
