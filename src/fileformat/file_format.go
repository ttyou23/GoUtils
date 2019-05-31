// file_format
package fileformat

import (
	"fmt"
	"os"
	"strings"
)

const DEFAULT_SWITCH_SIZE int64 = 1888
const ENCODE_TAG string = ".se"

func FileFormat(path string, isEncode bool) bool {

	if path == "0" {
		dir, err := GetCurrentPath()
		if err != nil {
			fmt.Println("GetCurrentPath err:", err)
			return false
		}
		path = dir + string(os.PathSeparator) + "encode"
	}

	fmt.Println("GetCurrentPath dir:", path)
	files, err := ListDir(path, ".exe", false)
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

	if strings.HasSuffix(strings.ToLower(file.filename), ENCODE_TAG) { //匹配文件
		return true
	}

	if interchangeFile(file.fullpath) {
		newFileName := Base256_encode_gbk(file.filename)
		newFileName = strings.ReplaceAll(fmt.Sprintf("%3d", index), " ", "0") + "_" + newFileName + ENCODE_TAG
		fmt.Println(newFileName)

		newFilePath := file.filepath + string(os.PathSeparator) + newFileName
		err := os.Rename(file.fullpath, newFilePath)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func decode_file(file encode_file_item, index int) bool {

	if !strings.HasSuffix(strings.ToLower(file.filename), ENCODE_TAG) { //匹配文件
		return true
	}

	if interchangeFile(file.fullpath) {
		newFileName := strings.ReplaceAll(file.filename, ENCODE_TAG, "")
		newFileName = string([]rune(newFileName)[4:])
		newFileName = Base256_decode_gbk(newFileName)
		fmt.Println(newFileName)

		newFilePath := file.filepath + string(os.PathSeparator) + newFileName
		err := os.Rename(file.fullpath, newFilePath)
		if err != nil {
			return false
		}
	} else {
		return false
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
