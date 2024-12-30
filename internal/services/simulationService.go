package service

import "project-golang/internal/repository"

type ISimulationService interface {
}

type SimulationService struct {
	repository *repository.IRepository
}

func NewSimulationService(repo *repository.IRepository) ISimulationService {
	return &SimulationService{repository: repo}
}

func CheckAntiFraude() bool {
	return false
}

func TokenJWTValido() bool {
	return false

}

func CalcularJuros() float64 {
	return 0.0
}

func CredorPossuiServicoSimulaçãoContratado() bool {
	return false

}
