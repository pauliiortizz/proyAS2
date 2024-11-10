package utils

import "net/http"

type HttpClient struct{}

type HttpClientInterface interface {
	Get(url string) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}

func (h *HttpClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}
