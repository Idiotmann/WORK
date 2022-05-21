package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..
	// _, err = conn.Do("Set", "name", "tomjerry 猫猫")   //string
	// _, err = conn.Do("HSet", "user01", "name", "john")    //hash
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19) //hash多输入
	if err != nil {
		// fmt.Println("set err=", err)
		fmt.Println("hset err=", err)
		return
	}
	// _, err = conn.Do("HSet", "user01", "age", 18)
	// if err != nil {
	// 	// fmt.Println("set err=", err)
	// 	fmt.Println("hset err=", err)
	// 	return
	// }
	// r, err := redis.String(conn.Do("Get", "name"))
	// r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	// if err != nil {
	// 	fmt.Println("hget err=", err)
	// 	return
	// }
	//因为返回r 是interface{}
	//因为name 对应的值是string ,因此我们需要转换
	//nameString := r.(string)
	// fmt.Println("操作ok ", r)
	// fmt.Printf("操作ok r1=%v r2=%v \n", r1, r2)
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
