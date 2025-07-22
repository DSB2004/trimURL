package models

import (
	"context"
	"encoding/json"
	"log"
	"url_shortener/pkg/config"

	"github.com/google/uuid"
)

type URL struct {
	Id       uuid.UUID `json:"id"`
	Original      string `json:"url"`
	Method   []string `json:"method"`
	ExpireIn string   `json:"expireIn"`
}


func (url *URL) Save(ctx context.Context)(string,error){
	redis := config.GetRedis()
	data, err := json.Marshal(url)
	if err != nil {
		return "", err
	}
	if(url.Id==uuid.Nil){
		url.Id=uuid.New()
	}

	err = redis.Set(ctx, url.Id.String(), data, 0).Err()
	if err!=nil{
		return "", err 
	}
	return url.Id.String(),nil
}

func (url *URL) Get(ctx context.Context,id string)(error){
	redis := config.GetRedis()
	data, err := redis.Get(ctx, id).Result()

	if err!=nil{
		return err
	}

	if err := json.Unmarshal([]byte(data), &url); err != nil {
		log.Println("Failed to unmarshal URL from Redis:", err)
		return err
	}
	return nil
	
}