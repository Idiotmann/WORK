package main

import (
	"WORK/wechat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据..")
	// conn read在conn没有被关闭时，才会阻塞
	// 如果客户端关闭了conn就不会阻塞 读不到东西会报错
	_, err = conn.Read(buf[0:4])
	if err != nil {
		// fmt.Println("conn.Read()读取错误err=", err)
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
	fmt.Println("客户端输入的内容是：", mes)
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// var data []byte  参数已经有了就不用在声明

	// 先发送一个长度给对方
	pkgLen := uint32(len(data))
	// 再转成[]byte切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	_, err = conn.Write(buf[0:4]) //切片要写后面的
	if err != nil {
		fmt.Println("conn.Write发送失败err=", err)
		return
	}
	n, err := conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data)发送失败err=", err)
		return
	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&loginMes)反序列化失败err=", err)
		return
	}
	//判断登录用户是否和数据库中的用户信息一样
	// 判断完合法不合法 之后构建回复消息 调用回复消息函数
	var resMes message.Message
	resMes.Type = message.LoginMesType

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200 //200表示用户匹配
	} else {
		// 不合法
		loginResMes.Code = 500 //500状态码表示该用户不存在
		loginResMes.Error = "该用户不存在,请注册后使用.."
	}
	//现在loginResMes是结构体 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes)序列化失败err=", err)
		return
	}
	// 现在data是loginResMes序列化后，把它给resMes message.Message消息结构体之后传出去
	resMes.Data = string(data)
	// 对resMes序列化后发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes)失败err=", err)
		return
	}
	// 封装writePkg
	err = writePkg(conn, data)
	// if err != nil {
	// 	fmt.Println("writePkg(conn,data)失败err=", err)
	// 	return
	// }
	return
}

// 判断消息的类型   执行相应的处理函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录消息 serverProcessLogin
		err = serverProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("serverProcessLogin(conn,mes)失败err=", err)
			return
		}
	case message.LoginResMesType:
		// 处理返回消息
	case message.RegisterMesType:
		//处理注册消息
	default:
		fmt.Println("消息类型不正确无法处理..")
	}
	return
}

func process(conn net.Conn) {
	// 需要延时关闭
	defer conn.Close()
	// 循环读客户端发送的数据

	for {
		// 这里将读取数据包直接封装成一个函数readPkg() 返回Message,err
		// 读之前先开一个切片
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了连接,服务器端正常退出")
				return
			} else {
				fmt.Println("readPkg(conn)读取数据包错误err=", err)
				return
			}
		}
		// fmt.Println("读取客户端的数据：", mes)

		// fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes(conn,&mes)失败err=", err)
			return
		}
	}
}
func main() {
	// 提示信息
	fmt.Println("服务区器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("监听错误。错误信息err=", err)
		return //跳出该函数
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept()接收出错。错误信息err=", err)
		}
		// 一旦连接，启动协程与客户端保持通信  去定义一个与客户端通信的函数process()
		go process(conn)

	}
}
