package commons

import (
	"encoding/json"
)

func StructToMap(obj interface{}) map[string]interface{} {
	jsonObject, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(jsonObject, &result)
	return result
}
