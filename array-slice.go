package main

import (
	"fmt"
	"reflect"
)

func main() {

	//getReflectType()
	//数组
	//var arr = [...]string{"11", 3: "a", 5: "b"}
	//var slice = []string{"11", 3: "a", 5: "b"}
	//printArr(arr, slice)
}

//函数
func getReflectType() {
	a := 1
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())
}

//函数参数中
//数组
//切片
func printArr(arr [6]string, slice []string) {
	for _, v := range arr {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("。。。。。。。\n")
	for i, len := 0, len(slice); i < len; i++ {
		fmt.Printf("%v\n", slice[i])
	}
}
