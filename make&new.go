package main

func main() {

	/*******************************************make三种引用类型map slice channel的使用*****************************************************/
	//map
	m := make(map[string]int)
	m["a"] = 1

	//slice
	lenth := 3 //长度
	//capacity:=5//容量 默认容量是和长度相等
	s := make([]string, lenth) //capacity
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	println(len(s), cap(s)) //3,3

	s = append(s, "d")
	println(len(s), cap(s)) //4,6

	s = append(s, "e")
	println(len(s), cap(s)) //5,6

	s = append(s, "f")
	println(len(s), cap(s)) //6,6

	s = append(s, "g")
	println(len(s), cap(s)) //7,12
	//channel

	println(m, s) //0xc420045e78 [3/5]0xc42007a000

}
