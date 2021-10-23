package main

import "fmt"

func A (a int){
	fmt.Println("defer print",a)
}
//a=0,返回2，b++操作的返回值
func B(a int)(b int){
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}

//a=0,返回1，因为b是局部变量，b++操作的局部变量
func C(a int) int{
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}
func main()  {
	var a, b int
	a = 1
	b = 2
	defer  A(a)//这里先注册到defer链表中，从stack copy a到堆中
	a +=1
	b +=1
	fmt.Println(a,b)// 打印完成后，调用A

	var c int
	d := B(c)
	fmt.Println(c,d)
}
