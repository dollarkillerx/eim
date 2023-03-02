package utils

import (
	"net/http"
	"time"
)

func GetDefaultHttpClient(timeout int) *http.Client {
	if timeout == 0 {
		timeout = 3
	}
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}
