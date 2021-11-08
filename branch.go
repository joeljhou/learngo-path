package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
	"strconv"
)

//if
func ifEval(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func ifFileRead() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func switchGrade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("Wrong score：%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

//for
func forEval() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forConverToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//while
func whileEval() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//函数 两个返回值
func evalDiv(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

//两个返回值只接收使用一个
func evalCalc(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a + b
	case "*":
		result = a * b
	case "/":
		q, _ := evalDiv(a, b)
		result = q
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//Go语言指针不能运算
func zhizhen() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

//值传递
func swap(x, y int) int {
	var temp int
	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

//引用传递
func swap2(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	//if
	fmt.Println(ifEval(-1))
	fmt.Println(ifEval(89))
	fmt.Println(ifEval(111))

	//switch
	fmt.Println(switchGrade(77))
	fmt.Println(switchGrade(90))
	fmt.Println(switchGrade(59))

	//for
	forEval()
	fmt.Println(forConverToBin(5), forConverToBin(13))

	//while
	whileEval()

	//返回两个参数
	q, r := evalDiv(13, 3)
	fmt.Println(q, r)

	//函数
	fmt.Println(evalCalc(78, 2, "-"))
	fmt.Println(evalCalc(78, 2, "+"))
	fmt.Println(evalCalc(78, 2, "*"))
	fmt.Println(evalCalc(78, 2, "/"))

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	zhizhen()
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
}
