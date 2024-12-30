package apiantifraude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"project-golang/internal/domain/dto"
	"time"

	"github.com/sony/gobreaker"
)

func CheckAntiFraud(request *dto.AntiFraudRequest) (*dto.AntiFraudResponse, error) {
	//CRIANDO CONFIG BREAKER
	settings := gobreaker.Settings{
		Name: "AntifraudeAPI",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Abrir o circuito se mais de 5 falhas ocorrerem
			return counts.ConsecutiveFailures > 5
		},
		Timeout: 10 * time.Second, // Tempo antes de tentar reabrir o circuito
	}
	cb := gobreaker.NewCircuitBreaker(settings)

	reqFunc := func() (interface{}, error) {
		return antiFraudRequest(request)
	}

	result, err := cb.Execute(reqFunc)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar a API antifraude: %w", err)
	}
	response, ok := result.(dto.AntiFraudResponse)
	if !ok {

		return nil, fmt.Errorf("erro inesperado na resposta da API antifraude")
	}
	return &response, nil

}

func antiFraudRequest(request *dto.AntiFraudRequest) (*dto.AntiFraudResponse, error) {
	var response = &dto.AntiFraudResponse{}
	timeout := 5 * time.Second
	jsonData, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	req, err := http.NewRequest("POST", os.Getenv("API_ANTIFRAUD"), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Configurar o cliente HTTP com timeout
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição para a API antifraude: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inválida da API antifraude: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
