package main

func main() {
	//定义map
	//1：内建函数 make
	//2：关键字 Map

	var m map[string]int
	m=make(map[string]int)//不初始化，会创建一个nil map，nil map 不能用来存放键值对
	m["a"]=1
	m["b"]=2
}
