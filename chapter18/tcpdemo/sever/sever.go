package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环接受发送的数据
	defer conn.Close() //关闭
	fmt.Printf("\n\n客户端ip=%v已连接\n服务器在等待客户端的输入...\n", conn.RemoteAddr().String())
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// fmt.Printf("\n\n客户端ip=%v已连接\n服务器在等待客户端的输入...\n", conn.RemoteAddr().String())
		//conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn中读取 如果不发会一直堵塞
		if err != nil {
			fmt.Printf("客户端ip=%v退出\n", conn.RemoteAddr().String())
			// 读取到一个错误就不要再等待了
			return //!!!!!
		}
		// 显示终端发送的数据
		fmt.Printf("客户端ip=%v输入的数据为:%v\n", conn.RemoteAddr().String(), string(buf[:n])) //读取切片中的从0到N
	}

}
func main() {
	fmt.Println("服务器监听开始...")
	// 监听函数 返回listen接口 错误返回err
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("服务器监听不到接口,信息err=%v", err)
		// 监听不到这个接口就干掉，就是没这个接口
		return //!!!!!
	}
	// 监听到这接口 要保证监听不能走，一直监听
	defer listen.Close() //延时关闭listen

	// 循环等待客户连接我
	for {
		fmt.Println("等待客户端连接")
		// 连接返回的是con接口。在接口中定义了读写等操作
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("服务器监听接口时发生错误,信息err=%v", err)

			// 即使没连接，或者与a连接错误，但是不能干掉return  因为可能还有b
		} else {
			fmt.Printf("con suc=%v\n客户端ip=%v\n", conn, conn.RemoteAddr().String())
			// 返回地址addr是接口 调用接口的string()字符串格式地址
		}
		//连接成功后继续从开始等待连接
		// 现在还没有客户端，用talent监听百度的，测试本程序
		// 接受数据不要在主线程，不然会很拥堵 用协程
		go process(conn)
	}
	// fmt.Printf("Listener suc=%v\n",listen)
}
