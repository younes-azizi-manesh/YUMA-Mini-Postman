package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"postman/internal/model"
)

func SendRequest(req model.RequestPayload) (model.ResponsePayload, error) {
	start := time.Now()
	client := &http.Client{}

	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = bytes.NewBufferString(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, req.URL, bodyReader)
	if err != nil {
		return model.ResponsePayload{}, err
	}

	// Set headers
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return model.ResponsePayload{}, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	// Try to parse JSON for pretty printing
	var parsedBody interface{}
	if err := json.Unmarshal(respBody, &parsedBody); err != nil {
		parsedBody = string(respBody)
	}

	// Collect headers
	respHeaders := make(map[string]string)
	for k, v := range resp.Header {
		respHeaders[k] = v[0]
	}

	return model.ResponsePayload{
		Status:     resp.StatusCode,
		Headers:    respHeaders,
		Body:       parsedBody,
		DurationMs: time.Since(start).Milliseconds(),
	}, nil
}
