/**
  create by yy on 2019/11/15
*/

package go_helper

// convert bool to int
func BToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

// convert int to bool
func IToB(i int) bool {
	return i != 0
}


	/***************/
	/*             */
	/*             */
	/*   string    */
	/*             */
	/*             */
	/***************/


// 字符串可以用==和<进行比较；比较通过逐个字节比较完成的，因此比较的结果是字符串自然编码的顺序。

// convert []byte to string
func ByteToStr(b []byte) string {
	return string(b)
}

// convert string to []byte
func StrToByte(s string) []byte {
	return []byte(s)
}
