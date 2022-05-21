package message

// import(
// 	_"fmt"
// )

// 确定消息的类型LoginMesType的名称就是LoginMes  LoginResMesType的名称就是LoginResMes
const (
	LoginMesType    = "LoginMes"    //登录消息类型
	LoginResMesType = "LoginResMes" //登录返回的消息类型
	RegisterMesType = "RegisterMes"
)

// 要进行序列化 给标签
type Message struct {
	Type string `json:"type"` //注意 点 是左上角的
	Data string `json:"data"` //消息的数据的类型
}

// 定义一个类型为LoginMes的结构体
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd "`
	UserName string `json:"userName"`
}
type LoginResMes struct {
	Code  int    `json:"code"`   //500 表示未注册 200表示登录成功
	Error string `json:"error "` //返回的错误信息
}
type RegisterMes struct {
}

// LoginMes   LoginResMes 都是给Message 中的Data
