package formatter

import (
	"encoding/json"
)

func GetStructFromInterface(data interface{} , structure interface{} ) (error) {
	bodyBytes, err := json.Marshal(data)
	if err !=nil {
		return err
	}
	err =json.Unmarshal(bodyBytes, &structure)
	return  err
}

