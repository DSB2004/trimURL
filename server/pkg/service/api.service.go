package service

import (
	"context"
	"time"
	"url_shortener/pkg/models"
)

func CreateShortURL(data *models.URLRequestBody)(string ,error) {
	c,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url models.URL
	url.Original = data.Url
	url.Method = data.Method
	url.ExpireIn = data.ExpireIn

	id, err := url.Save(c)

	if err != nil {
		return "",err
	}
	return id,nil
}

func GetShortURL(id string)(*models.URL,error){
	c,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url models.URL
	err:=url.Get(c,id)

	if err != nil {
		return nil,err
	}
	return &url,nil
}