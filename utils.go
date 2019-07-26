/**
  create by yy on 2019-07-26
*/

package go_helper

// 根据页码和 每页大小获取 对应的开始位置(用以sql的offset 或者 limit)
func GetStartPos(page int, size int) int {
	return (page - 1) * size
}
