package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	//slice本身是没有数据的，是对底层array的一个view
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6]) //2, 3, 4, 5
	fmt.Println("arr[:6] =", arr[:6])   //0, 1, 2, 3, 4, 5
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1) //2, 3, 4, 5, 6, 7
	s2 := arr[:]
	fmt.Println("arr[:] =", s2) //0, 1, 2, 3, 4, 5, 6, 7

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)  //100 3 4 5 6 7
	fmt.Println(arr) //0 1 100 3 4 5 6 7

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)  //100, 1, 2, 3, 4, 5, 6, 7
	fmt.Println(arr) //100, 1, 2, 3, 4, 5, 6, 7

	fmt.Println("Reslice")
	fmt.Println(s2) //100 1 100 3 4 5 6 7
	s2 = s2[:5]
	fmt.Println(s2) //100 1 100 3 4
	s2 = s2[2:]
	fmt.Println(s2) //100 3 4

	fmt.Println("Extending slice")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	//s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s2=[5 6], len(s2)=2, cap(s2)=3
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	//arr = 0 1 2 3 4 5 6 7
	//s2 = 5 6
	s3 := append(s2, 10) //5 6 10
	s4 := append(s3, 11) //5 6 10 11
	s5 := append(s4, 12) //5 6 10 11 12
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	fmt.Println("arr =", arr) //0 1 2 3 4 5 6 10
}
