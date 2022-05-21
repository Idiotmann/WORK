package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.137.1:8888")
	if err != nil {
		fmt.Println("Dial() err=", err)
		return
	}
	reader := bufio.NewReader(os.Stdin) //代表标准输入[终端]
	fmt.Println("\n请输入发送数据..发送exit退出客户端：")
	for {
		// fmt.Println("conn成功=", conn)
		// 功能一：客户端可以发送单行数据，然后就退出
		// fmt.Println("\n请输入发送数据..发送exit退出客户端：")
		// reader := bufio.NewReader(os.Stdin) //代表标准输入[终端]
		line, err := reader.ReadString('\n')
		// 再读取的时候带着/n，要去掉
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		// ReadString读取直到第一次遇到\n
		if err != nil {
			println("读取错误 err=", err)
		}
		// 去掉\n

		line = strings.Trim(line, "\r\n")
		// 将line发送给服务器
		if line == "exit" {
			break
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			println("发送错误 err=", err)
		} else {
			fmt.Printf("发送了%d字节的数据\n", n)
		}
	}
}
