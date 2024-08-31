package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiClientConfig struct {
	BaseURL string
	Token   string
	Timeout time.Duration
}

type ApiClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewApiClient(config ApiClientConfig) *ApiClient {
	return &ApiClient{
		baseURL: config.BaseURL,
		token:   config.Token,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func Request[T any](c *ApiClient, endpoint string, data interface{}) (*T, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка на этапе конвертации дата в json: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка во время формирования тела запроса: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("key", c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка во время выполения http-запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка, получен код ответа: %d", resp.StatusCode)
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ошибка при декодиравании ответа: %w", err)
	}

	return &result, nil
}
