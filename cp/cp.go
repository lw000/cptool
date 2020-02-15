package cp

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tuyue/tuyue_tools/cptool/utils"
)

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func CopyDir(srcPath string, destPath string) error {
	utils.CreateDir(destPath)

	// 检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		log.Println(err)
		return err
	} else {
		if !srcInfo.IsDir() {
			err := errors.New("srcPath不是一个正确的目录！")
			log.Println(err)
			return err
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		log.Println(err)
		return err
	} else {
		if !destInfo.IsDir() {
			err := errors.New("destInfo不是一个正确的目录！")
			log.Println(err)
			return err
		}
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path = strings.ReplaceAll(path, "\\", "/")
			destNewPath := strings.Replace(path, srcPath, destPath, 1)
			// log.Println("复制文件:" + path + " 到 " + destNewPath)
			copyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		log.Printf(err.Error())
	}
	return err
}

// 生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return
	}
	defer srcFile.Close()
	// 分割path目录
	destSplitPathDirs := strings.Split(dest, "/")

	// 检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := utils.PathExists(destSplitPath)
			if b == false {
				// log.Println("创建目录:" + destSplitPath)
				// 创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
