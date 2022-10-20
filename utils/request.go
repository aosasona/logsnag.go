package utils

import (
	"encoding/json"
	"errors"
	"strings"
)

type Request struct {
	Url     string
	Headers map[string]string
}

type Header struct {
	Key   string
	Value string
}

type Body interface{}

type RequestConfig struct {
	Body    Body
	Headers []Header
}

type SendConfig struct {
	Url string
	RequestConfig
}

func (r *Request) New(baseUrl string, headers []Header) *Request {
	r.Url = strings.Trim(baseUrl, "/")
	if len(headers) > 0 {
		r.Headers = make(map[string]string)
		for _, header := range headers {
			r.Headers[header.Key] = header.Value
		}
	}
	return r
}

func (r *Request) AddHeader(key string, value string) *Request {
	r.Headers[key] = value
	return r
}

func (r *Request) Send(config SendConfig) {

}

func (r *Request) Get(endpoint string, config RequestConfig) error {
	method := "GET"
	endpoint = strings.Trim(endpoint, "/")
	r.Url = r.Url + "/" + endpoint

	if len(config.Headers) > 0 {
		for _, header := range config.Headers {
			r.Headers[header.Key] = header.Value
		}
	}

	jsonBody, err := json.Marshal(config.Body)
	if err != nil {
		return errors.New("cannot marshal body to json")
	}

}
