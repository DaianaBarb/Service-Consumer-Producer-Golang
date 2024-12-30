package dto

type AntiFraudeRequest struct {
	TomadorID string  `json: idtomador`
	Valor     float64 `json: valor`
}

type AntiFraudeResponse struct {
}
