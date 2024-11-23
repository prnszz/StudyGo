// 常量声明, const定义的常量,在程序编译时就确定了,不能修改
package main

import "fmt"

//常量声明的时候必须赋值
//常量声明的时候如果没有赋值,默认和上一个常量值相同

const b string = "abc"
const c = "abc"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println(a, b, c)

	type Allergen int

	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)
	println(IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)
	fmt.Println(IgEggs | IgChocolate | IgShellfish)

	type ByteSize float64

	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)

	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

	const (
		Apple, Banana = iota + 1, iota + 2
		Cherimoya, Durian
		Elderberry, Fig
	)
	println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
}
