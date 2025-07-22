package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
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
	if url.ExpireIn != "" {
		expireTime, err := time.Parse("2006-01-02T15:04:05Z07:00", url.ExpireIn)
		if err != nil {
			return "", fmt.Errorf("invalid expiration time format: %v", err)
		}
		expireDuration := time.Until(expireTime)
		if expireDuration <= 0 {
			return "", fmt.Errorf("expiration time must be in the future")
		}

		err = redis.Set(ctx, url.Id.String(), data, expireDuration).Err()
	} else {
		err = redis.Set(ctx, url.Id.String(), data, 24*time.Hour).Err()
	}
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