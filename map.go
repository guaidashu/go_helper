/**
  create by yy on 2019/11/18
*/

package go_helper

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type JsonToMapValue struct {
	Value map[string]interface{}
}

/**
zh:
返回值说明
param 1: int, 返回值的类型，
		 1为 interface{}值类型，
		 2为 空(nil)
param 2: 如果param为 1，则此返回值是有效返回值
en:
The description of return values
param 1: int, return type，
		 1 is interface{}, type of value，
		 2 is null(nil)
param 2: if param is 1，the return value is a valid return value.
*/
func (j *JsonToMapValue) Get(key ...string) (int, interface{}) {
	var (
		length int
		err    interface{}
		tmp    map[string]interface{}
	)
	defer func() {
		if err = recover(); err != nil {
			switch err.(type) {
			case runtime.Error: // 运行时错误
				fmt.Println("runtime error:", err)
			default: // 非运行时错误
				fmt.Println("error:", err)
			}
		}
	}()

	length = len(key)

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

	tmp = j.Value

	// 现在处理 多层map的情况
	for k, v := range key {
		if k == length-1 {
			return 1, tmp[v]
		}
		tmp = tmp[key[k]].(map[string]interface{})
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
