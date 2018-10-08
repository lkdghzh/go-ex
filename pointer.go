package main

import "fmt"

func main() {
	//runPointer()

	//值类型
	var a int = 5
	var b int = 6
	//传地址
	funcArgsPointer(&a, b)
	fmt.Println(a, b) //10 6
}

//指针*  （定义 *类型）  （取值  *变量）
//取地址符&
func runPointer() {
	a := 10
	fmt.Println(&a) //0xc420045f70  &取地址  取变量的地址

	var p1 *int = &a
	fmt.Println(p1) //0xc420045f70

	p2 := &a        //b被自动推到类型成 *int
	fmt.Println(p2) //0xc420045f70

	//var b int
	//b=&a//cannot use &a (type *int) as type int in assignment
	//fmt.println(b)

	fmt.Println(*p1, p1) //取指针的值，*指针变量p1（这个指针变量是个地址）

	//*的作用
	//定义,在类型的前面加上*    var p *int
	//取值,在变量前加上*（变量是个地址，*变量是取这个地址对应的值） *p

}

//函数中参数指针的应用
func funcArgsPointer(m *int, n int) {
	*m = 10
	n = 20
}
