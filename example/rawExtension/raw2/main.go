package main

import (
	"encoding/json"
	"fmt"
)

// ServingInfo 结构体表示 servingInfo 字段的结构
type ServingInfo struct {
	CipherSuites  []string `json:"cipherSuites"`
	MinTLSVersion string   `json:"minTLSVersion"`
}

// TargetCSIConfig 结构体表示 targetcsiconfig 字段的结构
type TargetCSIConfig struct {
	ServingInfo ServingInfo `json:"servingInfo"`
}

// ObservedConfig 结构体表示 ObservedConfig 字段的结构
type ObservedConfig struct {
	TargetCSIConfig TargetCSIConfig `json:"targetCSIConfig"`
}

func main() {
	// 假设 rawExtension 是你的 runtime.RawExtension 对象
	rawExtension := []byte(`{
		"targetcsiconfig": {
			"servingInfo": {
				"cipherSuites": [
					"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
					"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
					"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
					"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
					"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
					"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"
				],
				"minTLSVersion": "VersionTLS12"
			}
		}
	}`)
	// 打印rawExtension
	fmt.Printf("rawExtension的值: %v", rawExtension)
	fmt.Printf("rawExtension的类型: %T", rawExtension)

	// 定义一个目标对象
	var result ObservedConfig

	// 使用 json.Unmarshal 解码 JSON 数据
	err := json.Unmarshal(rawExtension, &result)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}

	// 打印解码后的对象
	fmt.Printf("解码结果: %+v\n", result)
	// 打印CipherSuites 的值
	fmt.Printf("CipherSuites的结果: %+v\n", result.TargetCSIConfig.ServingInfo.CipherSuites)
	// 打印CipherSuites 的值
	fmt.Printf("minTLSVersion的结果: %+v\n", result.TargetCSIConfig.ServingInfo.MinTLSVersion)
}
