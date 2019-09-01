/**
  create by yy on 2019-08-31
*/

package go_helper

import (
	"fmt"
	"os"
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
// CreateMsPath("test.txt", true, true)
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

func createMsPath(path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return err
	}
	defer file.Close()
	return nil
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
