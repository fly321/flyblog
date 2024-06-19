package interfaces

type PhpArray interface {
	// ArrayMerge 合并数组
	ArrayMerge(array1, array2 map[string]interface{}) map[string]interface{}
	// ArrayToJson 数组转json
	ArrayToJson(array map[string]interface{}) string
	// JsonToArray json转数组
	JsonToArray(json string) map[string]interface{}
	// ArrayPush arrayPush
	ArrayPush(array map[string]interface{}, value interface{}) map[string]interface{}
}
