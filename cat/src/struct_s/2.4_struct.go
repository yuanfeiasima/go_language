package strcut_s

import f "fmt"

type person struct {
	name string
	age int
}

func ask (){
	var p person
	p.name = "astexis"
	p.age = 10
	f.Printf("name is :%s, age is :%d", p.name , p.age)
}
