// file_util
package fileformat

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type encode_file_item struct {
	filename string
	filepath string
	fullpath string
}

//获取程序当前目录
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string, is_suffix bool) (files []encode_file_item, err error) {
	files = make([]encode_file_item, 0)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}

		has_suffix := strings.HasSuffix(strings.ToUpper(fi.Name()), suffix)
		bool_suffix := (has_suffix && is_suffix) || (!has_suffix && !is_suffix)

		if bool_suffix { //匹配文件
			var file_item encode_file_item
			file_item.filename = fi.Name()
			file_item.filepath = dirPth
			file_item.fullpath = dirPth + PthSep + fi.Name()
			files = append(files, file_item)
		}
	}

	return files, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}
