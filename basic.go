package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//通过括号省略重复的var
var (
	aa = 3
	ss = "kkk"
	bb = true
)

//初始化
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//初始化赋值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//自动推断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//简写方式
func variableShorter() {
	//只能在函数内使用
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//欧拉公式
func euler() {
	//e ^ (i * π) + 1 = 0
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

//强制类型转换/三角形
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 //const数值可作为各种类型使用
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举(特殊常量)
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	//b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		pt
	)
	fmt.Println(cpp, "", python, golang, javascript)
	fmt.Println(b, kb, mb, gb, pt)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()

	consts()
	enums()
}
