package main

import (
	"fmt"

	"project-golang/internal/fx/server"
)

func main() {
	fmt.Println("iniciando aplicação")
	fmt.Println("server.Start")

	server.Start()

	fmt.Println("terminando aplicação")

}
