package main

import "sort"

func main() {
	/************************定义map的方式1：内建函数 make和2：关键字 Map*****************************/
	//定义1
	var m1 map[string]int = make(map[string]int) //不初始化，会创建一个nil map，nil map 不能用来存放键值对
	m1["a"] = 1
	m1["b"] = 2
	for k, v := range m1 {
		println(k, v)
	}
	//定义2
	var m2 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m2 {
		println(k, v)

		//d 4
		//a 1
		//b 2
		//c 3
	}

	/***********************************map是无顺序的****************************************/

	//排序map
	var sortedKeys []string = []string{}
	for k, _ := range m2 {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys) //导入sort package 按照字符串排序
	for _, k := range sortedKeys {
		println(m2[k])
	}
}
