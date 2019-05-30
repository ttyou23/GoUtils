// file_format
package fileutils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const DEFAULT_SWITCH_SIZE int64 = 1888
const ENCODE_TAG string = ".se"

type encode_file_item struct {
	filename string
	filepath string
	fullpath string
}

func FileFormat(path string, isEncode bool) bool {

	if path == "0" {
		dir, err := GetCurrentPath()
		if err != nil {
			fmt.Println("GetCurrentPath err:", err)
			return false
		}
		path = dir
	}

	fmt.Println("GetCurrentPath dir:", path)
	files, err := ListDir(path, ".exe")
	if err != nil {
		fmt.Println("ListDir err:", err)
		return false
	}
	for i, filepath := range files {

		if isEncode {
			encode_file(filepath, i)
		} else {
			decode_file(filepath, i)
		}
	}

	return true
}

func encode_file(file encode_file_item, index int) bool {

	fmt.Println("encode_file")
	if strings.HasSuffix(strings.ToLower(file.filename), ENCODE_TAG) { //匹配文件
		return true
	}

	newFileName := Base256_encode_gbk(file.filename)
	newFileName = strings.ReplaceAll(fmt.Sprintf("%3d", index), " ", "0") + "_" + newFileName + ENCODE_TAG
	fmt.Println(newFileName)

	newFilePath := file.filepath + string(os.PathSeparator) + newFileName
	err := os.Rename(file.fullpath, newFilePath)
	if err != nil {
		return false
	}
	return true
}

func decode_file(file encode_file_item, index int) bool {
	fmt.Println("decode_file")
	if !strings.HasSuffix(strings.ToLower(file.filename), ENCODE_TAG) { //匹配文件
		return true
	}
	return true
}

func interchangeFile(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Stat err:", err)
		return false
	}

	switch_size := DEFAULT_SWITCH_SIZE
	if fileInfo.Size() < 2*DEFAULT_SWITCH_SIZE {
		switch_size = (int64)(fileInfo.Size() / 2)
	}

	fmt.Println("switch_size: %d", switch_size)
	fp, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return false
	}
	defer fp.Close()

	off, err := fp.Seek(-switch_size, os.SEEK_END)
	if err != nil {
		fmt.Println("ReadAt SEEK_END Seek err:", err)
		return false
	}
	end_bytes := make([]byte, switch_size)
	count, err := fp.ReadAt(end_bytes, off)
	if err != nil {
		fmt.Println("ReadAt err:", err, count)
		return false
	}

	off, err = fp.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println("ReadAt SEEK_SET Seek err:", err)
		return false
	}
	head_bytes := make([]byte, switch_size)
	count, err = fp.ReadAt(head_bytes, off)
	if err != nil {
		fmt.Println("ReadAt err:", err)
		return false
	}

	off, err = fp.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println("WriteAt SEEK_SET Seek err:", err)
		return false
	}
	count, err = fp.WriteAt(end_bytes, off)
	if err != nil {
		fmt.Println("end_bytes WriteAt err:", err)
		return false
	}

	off, err = fp.Seek(-switch_size, os.SEEK_END)
	if err != nil {
		fmt.Println("WriteAt SEEK_END Seek err:", err)
		return false
	}
	count, err = fp.WriteAt(head_bytes, off)
	if err != nil {
		fmt.Println("head_bytes WriteAt err:", err)
		return false
	}

	fmt.Println("interchangeFile successful")
	return true
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
func ListDir(dirPth string, suffix string) (files []encode_file_item, err error) {
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
		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
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
