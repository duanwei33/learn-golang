package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

type MyObject struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

type MyParentObject struct {
	MyObject MyObject `json:"field3"`
}

func main() {
	rawData := []byte(`{"field3":{"field1": "value1", "field2": 42}}`)

	rawExtension := runtime.RawExtension{Raw: rawData}

	var result MyParentObject
	err := json.Unmarshal(rawExtension.Raw, &result)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	fmt.Printf("result解析结果: %+v\n", result)
	fmt.Printf("result.MyObject.Field2解析结果: %+v\n", result.MyObject.Field2)
}
