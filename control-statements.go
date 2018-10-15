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

	//break label一样
	for i := 0; i < 2; i++ {
		//loop2:
		for j := 0; j < 2; j++ {
			break //break loop2
			println("inner")
		}
		println("outer")
	}
l:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			println("inner") //只打印一次inner
			break l
		}
		println("outer")
	}

	//goto

	a := 1
t:
	if a < 10 {
		a++

		goto t //不主张使用goto语句， 以免造成程序流程的混乱
	} else {
		println(a) //10
	}
}
