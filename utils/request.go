package utils

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

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

func (r *Request) AppendHeader(key string, value string) *Request {
	r.Headers[key] = value
	return r
}

func (r *Request) Send(config SendConfig) (map[string]interface{}, error) {
	var (
		req         *http.Response
		res         map[string]interface{}
		contentType string
		err         error
	)

	url := makeEndpoint(r.Url, config.Url)

	if r.Headers["Content-Type"] != "" {
		contentType = r.Headers["Content-Type"]
	} else {
		contentType = "application/json"
	}

	client := &http.Client{}

	switch config.method {
	case GET:
		initialRequest, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		initialRequest.Header.Add("Content-Type", contentType)
		req, err = client.Do(initialRequest)
		break
	case POST:
		initialRequest, err := http.NewRequest("POST", url, config.Body.(io.Reader))
		if err != nil {
			return nil, err
		}
		for key, header := range r.Headers {
			initialRequest.Header.Add(key, header)
		}
		req, err = client.Do(initialRequest)
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("invalid method")
	}
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(req.Body)

	err = serializeResponseToStruct(req.Body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Request) Get(endpoint string, config RequestConfig) (map[string]interface{}, error) {
	var (
		err error
		res map[string]interface{}
	)

	method := GET

	if len(config.Headers) > 0 {
		for _, header := range config.Headers {
			r.AppendHeader(header.Key, header.Value)
		}
	}

	config.Body, err = serializeBodyToBuffer(config.Body)
	if err != nil {
		return nil, err
	}

	res, err = r.Send(SendConfig{
		Url:    endpoint,
		method: method,
		Body:   config.Body,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Request) Post(endpoint string, config RequestConfig) (map[string]interface{}, error) {
	var (
		err error
		res map[string]interface{}
	)

	method := POST

	if len(config.Headers) > 0 {
		for _, header := range config.Headers {
			r.AppendHeader(header.Key, header.Value)
		}
	}

	config.Body, err = serializeBodyToBuffer(config.Body)
	if err != nil {
		return nil, err
	}

	res, err = r.Send(SendConfig{
		Url:    endpoint,
		method: method,
		Body:   config.Body,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
