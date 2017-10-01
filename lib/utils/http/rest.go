package http

import (
	"io"
	"net/http"
	"time"
)

// RestClient :nodoc:
type RestClient struct {
	c *http.Client
}

// WithClient :nodoc:
func WithClient(httpClient *http.Client) *RestClient {
	return &RestClient{
		httpClient,
	}
}

// NewClient :nodoc:
func NewClient() *RestClient {
	return &RestClient{
		&http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// WithTemplate :nodoc:
// func WithTemplate(templateTag string, templateName string, data interface{}) (buf bytes.Buffer, err error) {
// 	asset, err := utils.Asset(templateTag)
// 	if err != nil {
// 		return
// 	}

// 	t := template.New(templateName)
// 	t, err = t.Parse(string(asset[:]))
// 	if err != nil {
// 		return
// 	}

// 	err = t.Execute(&buf, data)
// 	return
// }

// Post :nodoc:
func (c *RestClient) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return c.request("POST", url, body)
}

// Put :nodoc:
func (c *RestClient) Put(url string, body io.Reader) (resp *http.Response, err error) {
	return c.request("PUT", url, body)
}

func (c *RestClient) request(method, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	return c.c.Do(req)
}
