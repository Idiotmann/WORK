package main

import (
	"WORK/wechat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据..")
	// conn read在conn没有被关闭时，才会阻塞
	// 如果客户端关闭了conn就不会阻塞 读不到东西会报错
	_, err = conn.Read(buf[0:4])
	if err != nil {
		fmt.Println("conn.Read()读取错误err=", err)
		// err = errors.New("read pkg header error")
		return
	}
	// 转换
	pkgLen := binary.BigEndian.Uint32(buf[0:4])
	// 从conn中读到buf中
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen])读取信息错误err=", err)
		return
	}
	// 拿到 反序列化成message.Message 不用声明上面已经有参数
	// 函数体内的变修改函数体外的变量可以传入变量的地址&，
	// 函数内以指针的方式操作变量。
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes)反序列化失败err=", err)
		return
	}
	return
}

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	// var data []byte  参数已经有了就不用在声明

// 	// 先发送一个长度给对方
// 	pkgLen := uint32(len(data))
// 	// 再转成[]byte切片
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	_, err = conn.Write(buf[0:4]) //切片要写后面的
// 	if err != nil {
// 		fmt.Println("conn.Write发送失败err=", err)
// 		return
// 	}
// 	n, err := conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(data)发送失败err=", err)
// 		return
// 	}
// 	return
// }
