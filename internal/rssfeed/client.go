package rssfeed

import (
	"net/http"
	"time"
)

type client struct {
	httpClient http.Client
	timeout    time.Duration
}

func NewClient(timeout time.Duration) client {
	return client{httpClient: http.Client{}, timeout: timeout}
}
