package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo/cli/dto"
)

// TodoClient represents a client for fetching TODOs from a remote service.
type TodoClient struct {
	BaseURL string
}

// Fetch retrieves a Todo item by its ID from the configured TodoClient's BaseURL.
// It returns the Todo item and any error encountered during the API call or processing.
func (c TodoClient) Fetch(id int) (dto.Todo, error) {
	url := fmt.Sprintf("%s/todos/%d", c.BaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return dto.Todo{}, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.Todo{}, fmt.Errorf("error reading response body: %w", err)
	}

	var result dto.Todo
	if err := json.Unmarshal(body, &result); err != nil {
		return dto.Todo{}, fmt.Errorf("error unmarshalling response into Todo: %w", err)
	}

	return result, nil
}
