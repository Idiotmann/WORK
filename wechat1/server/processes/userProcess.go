package processes

import (
	"WORK/wechat1/common/message"
	"WORK/wechat1/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 创建结构体 UserProcess  考虑字段是方法中需要哪些东西
type UserProcess struct {
	Conn net.Conn
}

// 可以不用conn参数 结构体已经有了
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	// 分层模式 改动 tf调用下层的结构体赋值给tf

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	// err = writePkg(conn, data)
	err = tf.WritePkg(data)
	// if err != nil {
	// 	fmt.Println("writePkg(conn,data)失败err=", err)
	// 	return
	// }
	return
}
