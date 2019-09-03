package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
Call is a representation of a call, designed to cope for simple REST calls
*/
type Call struct {
	URL           string
	Method        string
	Params        map[string]string
	Headers       map[string]string
	Authorization interface{}
	ContentType   string
	Body          string
}

/*
Response is a simplified response for REST calls
*/
type Response struct {
	StatusCode int
	Body       []byte
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

	client := &http.Client{}

	if call.Method != "GET" && call.Method != "DELETE" {
		if call.ContentType != "" {
			r.Header.Set("Content-Type", call.ContentType)
		}
		r.Header.Add("Content-Length", strconv.Itoa(len(call.Body)))
	}

	resp, err := client.Do(r)

	response := Response{
		StatusCode: resp.StatusCode,
	}

	if resp.StatusCode != 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		response.Body = data
	}

	return &response, err
}
