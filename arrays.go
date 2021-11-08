package main

import "fmt"

//数组遍历
func arrayRange() {
	//遍历数组
	numbers := [6]int{1, 3, 2, 5, 8, 4}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d\t", numbers[i])
	}

	//range方式
	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println("\nnumbers的最大值和下表：", maxValue, maxi)

	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("numbers的和：", sum)
}

//打印数组
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray3(arr []int) {
	arr[1] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//数组定义
	var array1 [5]int
	array2 := [3]int{1, 3, 5}
	array3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool //四行五列
	fmt.Println(array1, array2, array3)
	fmt.Println(grid)

	arrayRange()

	//值传递
	printArray(array1)
	printArray(array3)
	fmt.Println(array1, array2, array3)

	//利用指针实现引用传递
	printArray2(&array1)
	printArray2(&array3)
	fmt.Println(array1, array2, array3)

	printArray3(array1[:])
	printArray3(array3[:])
	fmt.Println(array1, array2, array3)
}
