package models

import (
	"errors"
	"time"
)

type URLRequestBody struct {
	Url      string   `json:"url"`
	Method   []string `json:"method"`
	ExpireIn string   `json:"expireIn"`
}

func (body *URLRequestBody) Validate() (*URLRequestBody,error) {
	if len(body.Url) == 0 || len(body.Url) < 8 || body.Url[0:8] != "https://" {
		return  nil, errors.New("Provide a valid url")
	}
	if len(body.Method) == 0 {
		body.Method = append(body.Method, "GET")
	}

	_, err := time.Parse(time.RFC3339, body.ExpireIn)
	if err != nil {
		body.ExpireIn = time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	}

	return body,nil
}