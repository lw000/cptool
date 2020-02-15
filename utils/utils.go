package utils

import (
	"log"
	"os"
)

// 检测文件夹路径时候存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dir string) {
	ok, err := PathExists(dir)
	if err != nil {
		log.Panic(err)
	}

	if !ok {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
	}
}
