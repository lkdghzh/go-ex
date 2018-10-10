package main

import (
	"fmt"
)

type baser interface {
	play()
	say()
}
type Cat struct {
	name string
}

func (c Cat) play() {
	println("cat play")
}
func (c Cat) say() {
	println("cat say")
}
func (c Cat) hello() {
	println("cat Hello")
}

func getSomething() baser {
	return Cat{}
}

func do(sb baser) {
	sb.play()
	sb.say()
}

//接口片段类型别名
type IS []interface{}

func main() {
	//a := getSomething();
	//a.say();
	//cat := Cat{}
	//do(cat)
	//
	//fmt.Println()

	//list := make(IS, 3)
	//list[0] = 1
	//list[1] = "a"
	//list[2] = Cat{}
	//
	//if v, ok := list[0].(int); ok {
	//	println(v)
	//}
	//for _,el :=range list{
	//
	//}

	list2 := []interface{}{1, 'm', Cat{"cc"}}

	if v, ok := list2[0].(int); ok {
		println(v) //1
	}
	if v, ok := list2[2].(Cat); ok {
		fmt.Println(v) //{cc}
	}
}
