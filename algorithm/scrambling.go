/**
  create by yy on 2019/11/25
*/

package algorithm

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
}

// Fisher-Yates随机置乱算法
func FinsherYatesScrambling(arr interface{}) {

	// 利用反射 进行 解析
	value := reflect.ValueOf(arr)

	finsherYatesScrambling(value)

}

type Interface interface {
	DeepCopy() interface{}
}

func finsherYatesScrambling(value reflect.Value, tp ...reflect.Type) {

	var (
		valueType reflect.Kind
		num       int
		tmp       reflect.Value
	)

	// 判断变量类型，如果有传入值，则表示 是指针类型传过来的对应值
	if len(tp) > 0 {
		valueType = tp[0].Kind()
	} else {
		valueType = value.Kind()
	}

	switch valueType {

	case reflect.Slice:

		length := value.Len() - 1

		for i := length; i >= 0; i-- {

			num = rand.Intn(i + 1)

			tmp = value.Index(num)

			current := value.Index(i)

			cpy2 := reflect.New(current.Type()).Elem()

			cpy2.Set(current)

			cpy := reflect.New(tmp.Type()).Elem()

			cpy.Set(tmp)

			value.Index(num).Set(cpy2)

			value.Index(i).Set(cpy)

		}

	case reflect.Ptr:

		if value.IsNil() {
			return
		} else {
			finsherYatesScrambling(value.Elem(), value.Elem().Type())
		}

	default:
		fmt.Println("Please input a slice.")

	}

}
