package main

import (
	"fmt"

	"github.com/jcantonio/rest-client/rest"
)

func main() {
	postMe := rest.Call{
		Method: "GET",
		URL:    "http://localhost:5000/companies",
		Headers: map[string]string{
			"serverSecret": "serverOwnSecretXXX",
		},
	}
	resp, err := rest.Do(postMe)
	fmt.Println(resp, err)

}
