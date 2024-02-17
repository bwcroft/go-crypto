package api

import "net/http"

type ClientHttp struct {
	BaseUrl string
	Client  *http.Client
}

func NewHttpClient(baseUrl string) *ClientHttp {
	return &ClientHttp{
		BaseUrl: baseUrl,
		Client:  &http.Client{},
	}
}

func (c *ClientHttp) Get(path string) (*http.Response, error) {
	fullUrl := c.BaseUrl + path
	return c.Client.Get(fullUrl)
}
