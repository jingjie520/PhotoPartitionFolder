package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	imageTypes = ".jpg,.jpeg,.png,.gif" //图片文件类型

	pathSeparator = "\\" //路径分割符
	mapExt        map[string]bool
	maxImage      = 500 //文件夹文件上限

	mapFolder map[string]int //文件夹

	CurrentFolder string //当前目录
)

/*
初始化
*/
func init() {
	mapExt = make(map[string]bool)
	mapFolder = make(map[string]int)
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
	if isIimage(path) {
		fmt.Println("发现图片：", path) //list the file
		//转存图片
		saveToFolder(path)
	}
	return nil
}

func saveToFolder(fullFilename string) {
	//获取照片拍摄日期
	folder := getCreateTimeByFullFilename(fullFilename)

	savePath := CurrentFolder + pathSeparator + folder

	val, ok := mapFolder[folder]

	if !ok {
		//建立文件夹
		err := os.Mkdir(savePath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return
		}

		mapFolder[folder] = 0
	}

	if val%500 == 0 && val > 0 {
		//文件夹重命名
		os.Rename(savePath, savePath+"_"+strconv.Itoa(val/500))
		os.Mkdir(savePath, os.ModePerm)
	}

	newFileName := filepath.Base(fullFilename)
	saveFileName := savePath + pathSeparator + newFileName

	//移动文件
	err := os.Rename(fullFilename, saveFileName)

	if err != nil {
		fmt.Printf("\n err=%s# \n Path1=%s \n path2=%s \n", err, fullFilename, saveFileName)
		return
	}

	mapFolder[folder]++
}

func getCreateTimeByFullFilename(fullFilename string) string {
	return "2018-10-26"
}

func main() {
	//获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前目录错误。", err)
		return
	}

	fmt.Println("正在获取当前文件夹下图片数量 ...")

	CurrentFolder = dir

	getFileList(CurrentFolder)

}

func getFileExt(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)
	return fileSuffix
}
