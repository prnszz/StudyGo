package main // 声明包名

import "fmt" // 引入包

func main() {
	/* 简单的程序 万能的hello world */
	fmt.Println("Hello Go")
}

/*
$ go run hello.go
-> Hello Go
*/

/*
 $go build test1_hello.go
 $./test1_hello
 Hello Go
*/
