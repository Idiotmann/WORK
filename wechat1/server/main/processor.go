package main

import (
	"WORK/wechat1/common/message"
	"WORK/wechat1/common/utils"
	"WORK/wechat1/server/processes"

	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 判断消息的类型   执行相应的处理函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录消息 serverProcessLogin
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
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

func (this *Processor) process() (err error) {
	for {
		// 这里将读取数据包直接封装成一个函数readPkg() 返回Message,err
		// 读之前先开一个切片
		// 创建Transfer 引用
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了连接,服务器端正常退出")
				return err
			} else {
				fmt.Println("readPkg(conn)读取数据包错误err=", err)
				return err
			}
		}
		// fmt.Println("读取客户端的数据：", mes)

		// fmt.Println("mes=", mes)
		// 创建引用
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes(conn,&mes)失败err=", err)
			return err
		}
	}
}
