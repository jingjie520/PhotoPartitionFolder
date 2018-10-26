package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	listfile   []string        //获取文件列表
	imageTypes = "jpg,png,gif" //图片文件类型

	pathSeparator = "\\" //路径分割符
	mapExt        map[string]bool
)

/*
初始化
*/
func init() {
	mapExt = make(map[string]bool)
	ostype := os.Getenv("GOOS") // 获取系统类型

	if ostype == "linux" {
		pathSeparator = "/" //Linux分隔符
	}

	imageExts := strings.Split(imageTypes, ",")
	for _, val := range imageExts {
		mapExt[val] = true
	}

}

/*
是否图片
*/
func isIimage(fileName string) bool {
	ext := getFileExt(fileName)

	fmt.Printf("fileName=%s,ext=%s\n", fileName, ext)

	_, ok := mapExt[ext]
	return ok
}

func getFileList(path string) string {
	err := filepath.Walk(path, Listfunc)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return " "
}

func Listfunc(path string, f os.FileInfo, err error) error {
	var strRet string
	strRet, _ = os.Getwd()

	strRet += pathSeparator

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path //+ "\r\n"

	if isIimage(strRet) {
		listfile = append(listfile, strRet) //将目录push到listfile []string中
		fmt.Println("发现图片：", strRet)        //list the file
	}
	return nil

}

func ListFileFunc(p []string) {
	for index, value := range p {
		fmt.Println("Index = ", index, "Value = ", value)
	}

}

func main() {
	//获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前目录错误。", err)
		return
	}

	fmt.Println("正在获取当前文件夹下图片数量 ...")

	getFileList(dir)
	ListFileFunc(listfile)

}

func getFileExt(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)
	return fileSuffix
}
