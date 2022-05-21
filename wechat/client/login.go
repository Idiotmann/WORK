package main

import (
	"WORK/wechat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 函数完成登录

func login(userId int, userPwd string) (err error) {
	// // 下一个就是开始定协议
	// fmt.Printf("userId = %d userPwd= %s", userId, userPwd)
	// return nil
	//1：连接到服务器端

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial()连接到服务器端失败err=", err)
		return
	}
	// 记得延时关闭
	defer conn.Close()
	//2： 准备通过conn发送消息  定义消息类型
	var mes message.Message
	mes.Type = message.LoginMesType
	// mes.Data=message.LoginResMesType
	// message中定义了消息的类型  现在实例化创建一个LoginMes LoginResMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// loginMes.UserName = userName
	// 3：序列化 返回byte类型切片 一般还要转换成string()
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes序列化失败err=", err)
		return
	}
	mes.Data = string(data)
	// 不能直接给到data 因为它是结构体不是string类型，要先序列化成string
	//再把mes序列化 得到要发送的数据[]byte
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes序列化失败err=", err)
		return
	}

	// 先发送长度，再发送消息本身 获取长度，转成[]byte
	// len返回的是int 转换成uint32类型
	pkgLen := uint32(len(data))
	// 再转成[]byte切片
	var buf [4]byte
	// Variables
	// var BigEndian bigEndian
	// 大端字节序的实现。
	// var LittleEndian littleEndian
	// 小端字节序的实现。
	// type ByteOrder ByteOrder就是上面的两个变量
	// type ByteOrder interface {   将后面的类型转换成前面的切片类型
	//     Uint16([]byte) uint16
	//     Uint32([]byte) uint32
	//     Uint64([]byte) uint64
	//     PutUint16([]byte, uint16)
	//     PutUint32([]byte, uint32)
	//     PutUint64([]byte, uint64)
	//     String() string
	// }
	// ByteOrder规定了如何将字节序列和 16、32或64比特的无符号整数互相转化。
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	_, err = conn.Write(buf[0:4]) //切片要写后面的
	if err != nil {
		fmt.Println("conn.Write发送失败err=", err)
		return
	}
	// var LoginResMes message.LoginResMes
	// LoginResMes.Code
	// fmt.Printf("消息长度发送成功=%d 内容是:%s", len(data), string(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data)发送失败err=", err)
		return
	}

	// 这里还要处理服务器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn)出错。错误信息err=", err)
	}

	// 解包 将mes.Data反序列化到LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	// 休眠20
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s")
	return
}
