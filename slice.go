/**
  create by yy on 2019-07-26
*/

package go_helper

// 判断一个 元素 是否存在数组(切片中)
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Prepend(sc *[]interface{}, value interface{}) *[]interface{} {
	result := make([]interface{}, 1)
	result[0] = value
	for _, v := range *sc {
		result = append(result, v)
	}
	return &result
}
