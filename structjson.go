package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	//声明一个结构体
	p := Person{12, "like"}

	//结构体标签的使用(解码生成标签中的字段名,编码匹配标签中的字段名)
	str, _ := json.Marshal(p)
	fmt.Printf("%s\n", str)

	/********************************************************/
	//结构体加标签
	strs := `{"id":"a","name":"lk"}`

	var u Person
	//fmt.Println(u,&u)
	////解码为结构体
	json.Unmarshal([]byte(strs), &u)

	fmt.Println(u)

	a, _ := json.Marshal(u)
	fmt.Println("%s\n", a)

}
