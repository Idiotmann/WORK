package main

import (
	"fmt"
	"os"
)

// 定义全局变量
var userId int
var userPwd string

func main() {
	// 接收用户的选择
	var key string
	// 判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登录聊天室 ")
		fmt.Println("\t\t\t 2 注册用户 ")
		fmt.Println("\t\t\t 3 退出系统 ")
		fmt.Println("请选择(1-3): ")

		fmt.Scanln(&key)
		//判断key值
		switch key {
		case "1":
			fmt.Println("欢迎登录聊天系统..")
			loop = false
		case "2":
			fmt.Println("开始注册用户..")
			loop = false
		case "3":
			fmt.Println("确定退出聊天系统？")
			// loop = false  //退出系统直接os.Exit(0)
			os.Exit(0)
		default:
			fmt.Println("输入错误,请输入1-3")
		}
		// 根据用户的输入显示新的菜单
		if key == "1" {
			// 用户登录
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			// 登录函数写到另一个文件login.go
			// 在同一个文件夹下的函数可以调用
			login(userId, userPwd)
			// if err != nil {
			// 	fmt.Println("登录失败..")
			// } else {
			// 	fmt.Println("登录成功..")
			// }
		} else if key == "2" {
			fmt.Println("注册..")
		}
	}

}
