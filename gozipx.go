package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 获取当前时间并格式化为年月日时分秒
	currentTime := time.Now().Format("20060102_150405")

	// 文件路径
	srcFile := "rtthread.bin"
	// 新的文件名
	newFileName := fmt.Sprintf("bms_fw_%s.bin", currentTime)

	// 重命名文件
	err := os.Rename(srcFile, newFileName)
	if err != nil {
		fmt.Printf("文件重命名失败: %v\n", err)
		return
	}
	fmt.Printf("文件已重命名为: %s\n", newFileName)

	// 创建一个新的zip文件
	zipFileName := newFileName + ".zip"
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		fmt.Printf("创建zip文件失败: %v\n", err)
		return
	}
	defer zipFile.Close()

	// 创建一个zip.Writer对象
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 打开源文件
	fileToZip, err := os.Open(newFileName)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer fileToZip.Close()

	// 创建压缩文件条目
	zipEntry, err := zipWriter.Create(filepath.Base(newFileName))
	if err != nil {
		fmt.Printf("创建zip条目失败: %v\n", err)
		return
	}

	// 将源文件内容复制到zip条目
	_, err = io.Copy(zipEntry, fileToZip)
	if err != nil {
		fmt.Printf("文件压缩失败: %v\n", err)
		return
	}

	fmt.Printf("文件已成功压缩为: %s\n", zipFileName)
}
