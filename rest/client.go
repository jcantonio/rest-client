package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	MethodGet                = "GET"
	MethodPost               = "POST"
	MethodPut                = "PUT"
	MethodDelete             = "DELETE"
	MethodPatch              = "PATCH"
	AuthorizationTypeToken   = "bearer-token"
	AuthorizationTypeXApiKey = "x-api-key"
)

/*
Call is a representation of a call, designed to cope for simple REST calls
*/
type Call struct {
	URL           string
	Method        string
	Params        map[string]string
	Headers       map[string]string
	Authorization Authorization
	ContentType   string
	Body          string
}

type Authorization struct {
	Type    string
	Content interface{}
}
type TokenContent string

/*
Response is a simplified response for REST calls
*/
type Response struct {
	StatusCode int
	Body       []byte
}

func NewCall(url, method string) Call {
	return Call{
		URL:     url,
		Method:  method,
		Headers: map[string]string{},
		Params:  map[string]string{},
	}
}

/*
Do executes a REST
*/
func Do(call Call) (*Response, error) {
	var url bytes.Buffer
	url.WriteString(call.URL)

	if call.Params != nil {
		url.WriteByte('?')
		for k, v := range call.Params {
			url.WriteString(k)
			url.WriteByte(',')
			url.WriteString(v)
		}
	}

	r, err := http.NewRequest(call.Method, url.String(), strings.NewReader(call.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range call.Headers {
		r.Header.Set(k, v)
	}
	switch call.Authorization.Type {
	case AuthorizationTypeToken:
		r.Header.Set("Authorization", fmt.Sprintf("%v", call.Authorization.Content))
	case AuthorizationTypeXApiKey:
		r.Header.Set("x-api-key", fmt.Sprintf("%v", call.Authorization.Content))
	}

	type Authorization struct {
		Type    string
		Content interface{}
	}
	type TokenContent string

	client := &http.Client{}

	if call.Method != MethodGet && call.Method != MethodDelete {
		if call.ContentType != "" {
			r.Header.Set("Content-Type", call.ContentType)
		}
		r.Header.Add("Content-Length", strconv.Itoa(len(call.Body)))
	}

	resp, err := client.Do(r)
	response := Response{}

	if err != nil {
		return nil, err
	}
	response.StatusCode = resp.StatusCode

	data, _ := ioutil.ReadAll(resp.Body)
	response.Body = data

	return &response, err
}
