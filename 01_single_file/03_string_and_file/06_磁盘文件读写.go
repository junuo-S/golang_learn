package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	// 打开文件或创建文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，关闭文件
	defer file.Close()
	for i := 0; i < 5; i++ {
		buf := fmt.Sprintf("i = %d\n", i)
		file.WriteString(buf)
	}
}

func ReadFile(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，关闭文件
	defer file.Close()
	// 读文件
	buf := make([]byte, 1024*2) // 2k大小
	n, err1 := file.Read(buf)
	if err1 != nil && io.EOF != err1 { // 文件出错，同时没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}

// 逐行读取
func ReadFileLine(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file.Close()
	// 创建一个 *Reader
	var reader = bufio.NewReader(file)
	for {
		// 遇到'\n'结束读取，但是'\n'也读取进去
		var buf, err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Println("buf = ", string(buf))
	}
}

func main() {
	var path = "test.txt"
	// WriteFile(path)
	// ReadFile(path)
	ReadFileLine(path)
}
