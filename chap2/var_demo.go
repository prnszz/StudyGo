package main

import "fmt"

var x, y int
var ( //这种分解的写法,一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "liudanbing"

//这种不带声明格式的只能在函数体内声明
//g, h := 123, "需要在func函数体内实现"

func main() {
	g, h := 123, "需要在func函数体内实现"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	//不能对g变量再次做初始化声明
	//g := 400

	_, value := 7, 5 //实际上7的赋值被废弃，变量 _  不具备读特性
	//fmt.Println(_) //_变量的是读不出来的
	fmt.Println(value) //5

	//第一种 使用默认值
	var a int
	fmt.Printf("a = %d\n", a)

	//第二种
	var b int = 10
	fmt.Printf("b = %d\n", b)

	//第三种 省略后面的数据类型,自动匹配类型
	var c = 20
	fmt.Printf("c = %d\n", c)

	//第四种 省略var关键字
	d := 3.14
	fmt.Printf("d = %f\n", d)
}
