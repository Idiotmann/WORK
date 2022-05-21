package utils

import (
	"WORK/wechat1/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 将这些方法联系到结构体中
type Transfer struct {
	// 分析 字段   1 conn buf
	Conn net.Conn
	Buf  [8096]byte //这时传输时用到缓冲

}

// 将函数改成方法  Conn不用给 用结构体里面的

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据..")
	// conn read在conn没有被关闭时，才会阻塞
	// 如果客户端关闭了conn就不会阻塞 读不到东西会报错
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		// fmt.Println("conn.Read()读取错误err=", err)
		// err = errors.New("read pkg header error")
		return
	}
	// 转换
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])
	// 从conn中读到buf中
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen])读取信息错误err=", err)
		return
	}
	// 拿到 反序列化成message.Message 不用声明上面已经有参数
	// 函数体内的变修改函数体外的变量可以传入变量的地址&，
	// 函数内以指针的方式操作变量。
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes)反序列化失败err=", err)
		return
	}
	fmt.Println("客户端输入的内容是：", mes)
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// var data []byte  参数已经有了就不用在声明

	// 先发送一个长度给对方
	pkgLen := uint32(len(data))
	// 再转成[]byte切片
	// var buf [4]byte 只是作为缓冲使用 可以可上面的共用
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	_, err = this.Conn.Write(this.Buf[0:4]) //切片要写后面的
	if err != nil {
		fmt.Println("conn.Write发送失败err=", err)
		return
	}
	n, err := this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data)发送失败err=", err)
		return
	}
	return
}
