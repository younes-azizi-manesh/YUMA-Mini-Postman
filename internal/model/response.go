package model

type ResponsePayload struct {
	Status     int               `json:"status"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
	DurationMs int64             `json:"duration_ms"`
}
