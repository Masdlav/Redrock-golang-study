package main

//练习结构体的使用

import (
	"fmt"
)

type animal interface {
	say()
}

type Duck struct {
	Name string
	Color string
	Age int
}

func (Duck)say() {
	fmt.Println("嘎嘎嘎")
}

type Dog struct {
	Name string
	Color string
	Age int
}
func (Dog)say()  {
	fmt.Println("汪汪汪")
}

func hit(a animal)  {
	a.say()
}

func main() {
	daHuang := Dog{
		Name: "大黄",
		Color: "yellow",
		Age: 3,
	}
	// fmt.Println(daHuang.Name) 如果没有初始化string类型的属性则默认值为空
	keDa := Duck{
		Name: "可达鸭",
		Color: "yellow",
		Age: 2,
	}
	//fmt.Println(keDa.Age) 如果没有初始化int类型的属性则默认值为0
	hit(keDa)
	hit(daHuang)
}

