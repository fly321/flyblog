package test

import (
	"flyblog/src/helper"
	"fmt"
	"testing"
)

func TestPhpArray(t *testing.T) {
	merge := helper.PhpArray{}.ArrayMerge(map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2})
	fmt.Println(merge)

	json := helper.PhpArray{}.ArrayToJson(map[string]interface{}{"a": 1})
	fmt.Println(json)

	array := helper.PhpArray{}.JsonToArray(json)
	fmt.Println(array)

	push := helper.PhpArray{}.ArrayPush(map[string]interface{}{"a": 1}, 2)
	fmt.Println(push)

}
