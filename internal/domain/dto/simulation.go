package dto

type AntiFraudRequest struct {
	BorrowerId string  `json:"borrowerId"`
	Valor      float64 `json:"valor"`
}

type AntiFraudResponse struct {
	BorrowerId string `json:"borrowerId"`
}

type QueuePublishPayload struct {
	SimulationId string `json:"simulationId "`
}

type SimulationRequest struct {
}

type SimulationResponse struct {
}
