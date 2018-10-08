package main

//oop中的继承->(匿名字段)
//函数和方法的区别->（多一个接受者[值/指针接受者]）
//结构的方法的定义，可以通过接受者reciver【函数签名】来表示

//继承
type Person struct {
	name string
}
type GoSoftEngineer struct {
	Person        //匿名字段实现继承
	name   string //和父结构相同字段
	tech   string
}
type JSSoftEngineer struct {
	Person //匿名字段实现继承
	tech   string
}

//方法不是函数method not func
//通过接受者可以取到其余字段，不同struct的method不同
//当父结构中也同字段，通过匿名字段获取父struct的字段
func (goEngineer GoSoftEngineer) coding(code string) string {
	return "原名" +
		goEngineer.Person.name + "，人称之" +
		goEngineer.name + "," +
		goEngineer.tech + "," +
		code + ",go"
}
func (jsEngineer JSSoftEngineer) coding(code string) string {
	return "原名" +
		jsEngineer.name + "," +
		jsEngineer.tech + "," +
		code + ",js"
}

//https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (goSoftEngineer *GoSoftEngineer) changeAlias() {
	goSoftEngineer.name = "统一改名了" //goSoftEngineer.name 底层go会调用 *goSoftEngineer.name 2（依照函数签名，就转成是值类型还是指针类型）
}
func (goSoftEngineer GoSoftEngineer) changeAlias2() {
	goSoftEngineer.name = "统一改名了2" //只是对副本做更改
}

func main() {
	lk := GoSoftEngineer{Person{"lkRD"}, "go高手", "goland"}
	lk2 := JSSoftEngineer{Person{"lkFE"}, "javascript"}
	println(lk.coding("lanuage"))  //原名lkRD，人称之go高手,goland,lanuage,go
	println(lk2.coding("lanuage")) //原名lkFE,javascript,lanuage,js

	println(lk.name)              //go高手
	lk.changeAlias2()             //副本非指针
	println(lk.name)              //go高手
	println(lk.coding("lanuage")) //原名lkRD，人称之go高手,goland,lanuage,go

	lk.changeAlias()              //这个地方应该	lk := &GoSoftEngineer{Person{"lkRD"}, "go高手", "goland"}，底层go做了转化，1（依照函数签名，就转成是值类型还是指针类型）
	println(lk.name)              //统一改名了
	println(lk.coding("lanuage")) //原名lkRD，人称之统一改名了,goland,lanuage,go

}
