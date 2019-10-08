/**
  create by yy on 2019-07-26
*/

package go_helper

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// 根据页码和 每页大小获取 对应的开始位置(用以sql的offset 或者 limit)
func GetStartPos(page int, size int) int {
	return (page - 1) * size
}

type JsonToMapValue struct {
	Value map[string]interface{}
}

/**
返回值说明
param 1: int, 返回值的类型，
		 1为 interface{}值类型，
		 2为 空(nil)
param 2: 如果param为 1，则此返回值是有效返回值
*/
func (j *JsonToMapValue) Get(key ...string) (int, interface{}) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error: // 运行时错误
				fmt.Println("runtime error:", err)
			default: // 非运行时错误
				fmt.Println("error:", err)
			}
		}
	}()
	length := len(key)
	if length < 1 {
		return 2, nil
	}

	// 如果传的是单值，则返回单值
	if length == 1 {
		// 首先获取一下key， 如果是存在的，则返回
		if v := j.Value[key[0]]; v != nil {
			return 1, v
		}

		return 2, nil
	}

	var tmp map[string]interface{}

	// 现在处理 多层map的情况
	for k, v := range key {
		if k == length-1 {
			return 1, tmp[v]
		}
		tmp = j.Value[key[k]].(map[string]interface{})
	}

	// 不存在，返回空
	return 2, nil
}

func JsonToMap(value []byte) *JsonToMapValue {
	var m map[string]interface{}
	err := json.Unmarshal(value, &m)
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return nil
	}
	return &JsonToMapValue{
		Value: m,
	}
}
