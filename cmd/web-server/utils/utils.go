package utils

import "encoding/json"

func ObjToJSONObj(obj interface{}) interface{} {
	var jsonMap interface{}
	b, _ := json.Marshal(obj)
	json.Unmarshal(b, &jsonMap)
	return jsonMap
}
