package main

import "fmt"

var a_var = "a是全局变量"
func main() {
	b_var := "b是局部变量"
	fmt.Println(a_var)
	fmt.Println(b_var)

	_, c := "会被下划线丢弃", "不会被下划线丢弃"
	//fmt.Println(_)
	fmt.Println(c)

	/*
	常量
	 */
	const constantName  = "go language"
	const pi float32 = 3.1415926
	const i = 100
	const MaxThread = 10
	const prefix = "prefix_str"

	fmt.Println(constantName)
	fmt.Println(pi)
	fmt.Println(MaxThread)
	fmt.Println(prefix)

	test()

}

/*
内置基础类型
boolean
数值类型
 */

/*
字符串
 */
var frenchHello string
var emptyString string = ""
func test(){
	no, yes, maybe := "no", "yes", "maybe"//简短声明， 同时声明多个变量
	japaneseHello := "Konichiwa" //简短声明
	frenchHello = "Bonjour" //常规赋值
	fmt.Println(no, yes, maybe)
	fmt.Println(japaneseHello)
	fmt.Println(frenchHello)

	s := "hello"
	//s[0] = "a"
	fmt.Println(s)

	c := []byte(s)
	c[0]= 'c'
	s_s := string(c)
	fmt.Println(s_s)

	m := "hello, "
	n := "world"
	l := m+n
	fmt.Println(l)

	s = "d"+ s[1:]//字符串虽然不能修改， 但是可以进行切片操作
	fmt.Printf("%s\n", s)


	//声明多行字符串
	mm := `duo hang
	zifuchuan`
	fmt.Println(mm)

}
