/**
  create by yy on 2019-08-31
*/

package go_helper

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// The second param (bool)
// This param is used to let you choose whether create Multistage path.
// example:
// If you given path which called "./files/test/test.txt",
// and "./files" or "./files/test" is not exists, the program will first create them.
//
// The third param (bool)
// If you want to create a directory
// Please given the third param,
// example:
// CreateMsPath("test", true, true)
func CreateMsPath(args ...interface{}) error {
	var err error
	path := args[0].(string)
	flag := true
	isCreateDirectory := false
	if len(args) > 1 {
		flag = args[1].(bool)
	}
	if len(args) > 2 {
		isCreateDirectory = args[2].(bool)
	}
	if !flag {
		if isCreateDirectory {
			err := createMsDir(path)
			return err
		}
		err := createMsPath(path)
		return err
	}
	if isCreateDirectory {
		err = createMsAllDir(path)
		return err
	}
	path = strings.Replace(path, "\\", "/", -1)
	pathArr := strings.Split(path, "/")
	tmpPath := ""
	pathLength := len(pathArr) - 1
	pathLengthSecond := pathLength - 1
	for k, item := range pathArr {
		tmpPath = getPath(tmpPath, item)
		if ok, _ := PathExists(tmpPath); ok {
			continue
		}
		if pathLength == k {
			err = createMsPath(tmpPath)
		} else if pathLengthSecond == k {
			err = createMsAllDir(tmpPath)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func getPath(path string, item string) string {
	if item == "." && path == "" {
		return "."
	}
	if path == "" {
		return item
	}
	return path + "/" + item
}

func createMsPath(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	defer func() {
		if err == nil {
			err = file.Close()
		} else {
			_ = file.Close()
		}
	}()
	return
}

func createMsDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func createMsAllDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

/**
读取文件， buffer方式，可以读大文件
param 1: file path
param 2: buffer size
*/
func ReadFile(params ...interface{}) string {
	var path string
	if len(params) > 0 {
		path = params[0].(string)
	} else {
		return ""
	}
	var size int
	if len(params) > 1 {
		size = params[0].(int)
	} else {
		size = 512
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return ""
	}
	s := ""
	reader := bufio.NewReader(file)
	buf := make([]byte, size)
	for {
		n, err := reader.Read(buf)
		s = s + string(buf[:n])
		if n < size || n <= 0 || err == io.EOF {
			break
		}
	}
	return s
}

/**
读取文件， buffer方式，可以读大文件
param 1: file path
param 2: buffer size
*/
func ReadFileToByte(params ...interface{}) []byte {
	var path string
	if len(params) > 0 {
		path = params[0].(string)
	} else {
		return nil
	}
	var size int
	if len(params) > 1 {
		size = params[0].(int)
	} else {
		size = 512
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil
	}
	var s []byte
	reader := bufio.NewReader(file)
	buf := make([]byte, size)
	for {
		n, err := reader.Read(buf)
		s = append(s, buf[:n]...)
		if n < size || n <= 0 || err == io.EOF {
			break
		}
	}
	return s
}

// 写内容到文件
// Write content to a file.
// If there is not exists, the method will auto create it.
func WriteStringToFile(path string, params ...interface{}) {
}

// 过滤给定字符串(这个函数的目的开始是为了过滤我的json配置文件里的自定义注释)里的 注释，根据输入的 注释开头 和 结尾进行过滤
// 过滤的 开头和结尾需要转移特殊字符， 具体方式看例子
// example:
//	FilterComment(data, "/\\*", "\\*/")
func FilterComment(originStr, startStr, endStr string) string {
	// 利用正则匹配替换
	patter := "(" + startStr + `[\w\W]*?` + endStr + `)`
	rg, _ := regexp.Compile(patter)
	s := rg.ReplaceAllString(originStr, "")
	return s
}
