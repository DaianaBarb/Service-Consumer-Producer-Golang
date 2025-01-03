package routes

import (
	"fmt"
	"project-golang/internal/api/rest/handlres"
	"time"

	"net/http"

	//_ "internal/docs"

	"github.com/gorilla/mux"
	// Swagger UI
	// Swagger middleware
)

type IRoutes interface {
	RegisterRoutes()
}

func (s *Routes) RegisterRoutes() {

	c := mux.NewRouter()
	c.HandleFunc("/v1/generateJwt", s.handlerSimulation.GenerateJWTw).Methods("GET")
	c.HandleFunc("/v1/simulation", s.handlerSimulation.CreatedSimulation).Methods("POST")
	c.HandleFunc("/v1/setup", s.handlerSimulation.CreatedSetup).Methods("POST")
	c.HandleFunc("/v1/borrower", s.handlerSimulation.CreatedBorrower).Methods("POST")
	c.HandleFunc("/v1/simulation/{id}", s.handlerSimulation.FindByIdSimulation).Methods("GET")
	c.HandleFunc("/v1/setup/{id}", s.handlerSimulation.FindByIdSetup).Methods("GET")
	c.HandleFunc("/v1/borrower/{id}", s.handlerSimulation.FindByIdBorrower).Methods("GET")
	c.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/"))))
	c.HandleFunc("/v1/setup", s.handlerSimulation.UpdateSetup).Methods("PUT")
	c.HandleFunc("/v1/simulation/{id}", s.handlerSimulation.UpdateSimulation).Methods("PUT")
	c.HandleFunc("/v1/simulation/borrower/{id}", s.handlerSimulation.BorrowerResponseToSimulation).Methods("PATCH")
	c.HandleFunc("/v1/simulation", s.handlerSimulation.FindSimulationsByParam).Methods("GET").Queries()
	c.HandleFunc("/helth/db", s.handlerSimulation.HealthCheckHandler).Methods("GET")
	c.HandleFunc("/health/app", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"UP"}`))
		if err != nil {
			fmt.Println("error %v", err)
		}
	}).Methods("GET")

	fmt.Println(" online na porta 8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      c,
		ReadTimeout:  10 * time.Second, // Tempo máximo para leitura do corpo da requisição
		WriteTimeout: 10 * time.Second, // Tempo máximo para escrever a resposta
		IdleTimeout:  15 * time.Second, // Tempo máximo para conexões ociosas
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("error %v", err)
	}
}

type Routes struct {
	handlerSimulation handlres.ISimulationHandler
}

func NewRoutes(handler handlres.ISimulationHandler) IRoutes {

	return &Routes{
		handlerSimulation: handler,
	}

}
