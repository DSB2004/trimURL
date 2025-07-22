package utils

import "encoding/json"

func BodyParser(data []byte, body any) (error) {
	err := json.Unmarshal(data,body)
	if(err!=nil) {
		return err
	}
	return nil
}