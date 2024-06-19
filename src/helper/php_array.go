package helper

import (
	"encoding/json"
	"fmt"
)

type PhpArray struct {
}

func (receiver PhpArray) ArrayMerge(array1, array2 map[string]interface{}) map[string]interface{} {
	for k, v := range array2 {
		array1[k] = v
	}
	return array1
}

// ArrayToJson 数组转json
func (receiver PhpArray) ArrayToJson(array map[string]interface{}) string {
	jsonData, err := json.Marshal(array)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return ""
	}
	return string(jsonData)
}

// JsonToArray json转数组
func (receiver PhpArray) JsonToArray(jsonStr string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return nil
	}
	return data
}

// ArrayPush arrayPush
func (receiver PhpArray) ArrayPush(array map[string]interface{}, value interface{}) map[string]interface{} {
	array[fmt.Sprintf("%d", len(array))] = value
	return array
}
