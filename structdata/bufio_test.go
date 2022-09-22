package structdata

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBufIo(t *testing.T) {
	// bufio.NewReader()
	// 一般用于网络或者文件 io 操作，可以提高读\写 效率
	// bufio 拷贝文件
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	filePath := "c:/code/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("http://www.baidu.com/ \n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func TestRandomIo(t *testing.T) {
	Myfile, err := os.Open("c:/code/golang.txt")
	if err != nil {
		fmt.Println("Error opening file!!!")
	}

	byteBuff := make([]byte, 11)
	totalLen, err := Myfile.Read(byteBuff)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("File Data: \n%s\n", string(byteBuff[:totalLen]))

	//We move file pointer 3 bytes before from current position.
	newPosition, err := Myfile.Seek(3, 0)
	if err != nil {
		fmt.Println(err)
	}

	byteBuff1 := make([]byte, 5)
	totalLen1, err1 := Myfile.Read(byteBuff1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("File Data from position %d is: \n%s\n", newPosition, string(byteBuff1[:totalLen1]))

	Myfile.Close()

}
