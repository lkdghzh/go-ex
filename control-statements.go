package main

func main() {

	//if
	//条件判断语句里面允许声明一个变量,
	//变量的作用域只能在该条件逻辑块内，其他地方就不起作用
	if a := 1; a == 1 {
		println(true)
	} else {
		println(false)
	}
	//undefined: a
	//println(a)
}
