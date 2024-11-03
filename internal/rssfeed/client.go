package rssfeed

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	timeout    time.Duration
}

func NewClient(timeout time.Duration) Client {
	return Client{httpClient: http.Client{}, timeout: timeout}
}
