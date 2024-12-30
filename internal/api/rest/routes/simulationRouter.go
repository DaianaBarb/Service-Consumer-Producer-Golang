package routes

import (
	"fmt"
	"project-golang/internal/api/rest/handlres"

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

	c.HandleFunc("/v1/simulation", s.handlerSimulation.CreatedSimulation).Methods("POST")
	c.HandleFunc("/v1/setup", s.handlerSimulation.CreatedSetup).Methods("POST")
	c.HandleFunc("/v1/borrower", s.handlerSimulation.CreatedBorrower).Methods("POST")
	c.HandleFunc("/v1/simulation/{id}", s.handlerSimulation.FindByIdSimulation).Methods("GET")
	c.HandleFunc("/v1/setup/{id}", s.handlerSimulation.FindByIdSetup).Methods("GET")
	c.HandleFunc("/v1/borrower/{id}", s.handlerSimulation.FindByIdBorrower).Methods("GET")
	c.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/"))))
	c.HandleFunc("/v1/simulation", s.handlerSimulation.UpdateSetup).Methods("PUT")
	c.HandleFunc("/v1/simulation", s.handlerSimulation.UpdateSimulationStatus).Methods("PUT")
	c.HandleFunc("/v1/simulation/response-borrower", s.handlerSimulation.SimulationResponseBorrower).Methods("POST")

	fmt.Println(" online na porta 8080")
	http.ListenAndServe(":8080", c)
}

type Routes struct {
	handlerSimulation handlres.ISimulationHandler
}

func NewRoutes(handler handlres.ISimulationHandler) IRoutes {

	return &Routes{
		handlerSimulation: handler,
	}

}
