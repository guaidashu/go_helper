/**
  create by yy on 2019-07-26
*/

package go_helper

import (
	"errors"
	"fmt"
	"runtime"
)

// 根据页码和 每页大小获取 对应的开始位置(用以sql的offset 或者 limit)
func GetStartPos(page int, size int) int {
	return (page - 1) * size
}

/**
用于报告错误行数和文件名在哪里，便于找bug
This func is used to report the error line and file name
so that we can find bug quickly.

一般在项目中应用的时候，应该配置一个全局的控制变量，并且打开注释代码块里的注释，
根据你的全局变量进行修改，以达到可以关闭的效果，否则是默认都会报告的
*/
func NewReportError(err error) error {
	// if !config.Config.App.DEBUG {
	//	return err
	// }
	_, fileName, line, _ := runtime.Caller(1)
	data := fmt.Sprintf("%v, report in: %v: in line %v", err, fileName, line)
	return errors.New(data)
}
