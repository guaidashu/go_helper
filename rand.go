/**
  create by yy on 2019-09-11
*/

package go_helper

import "math/rand"

func GetIntRandom(interval int) int {
	return rand.Intn(interval)
}

func GetFloatRandom32(interval int) float32 {
	return rand.Float32() * float32(interval)
}

func GetFloatRandom64(interval int) float64 {
	return rand.Float64() * float64(interval)
}
