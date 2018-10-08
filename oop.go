package main

func main() {
	lk := GoSoftEngineer{Person{"lkRD"}, "go高手", "goland", 4}
	lk2 := JSSoftEngineer{Person{"lkFE"}, "javascript", 4}
	println(lk.coding("lanuage"))  //原名lkRD，人称之go高手,goland,lanuage,go
	println(lk2.coding("lanuage")) //原名lkFE,javascript,lanuage,js
}

//oop中的继承->(匿名字段)
//函数和方法的区别->（多一个接受者[值/指针接受者]）
//结构的方法的定义，可以通过接受者来表示

//继承
type Person struct {
	name string
}
type GoSoftEngineer struct {
	Person            //匿名字段实现继承
	name       string //和父结构相同字段
	tech       string
	experience int
}
type JSSoftEngineer struct {
	Person     //匿名字段实现继承
	tech       string
	experience int
}

//方法不是函数method not func
//通过接受者可以取到其余字段，不同struct的method不同
//当父结构中也同字段，通过匿名字段获取父struct的字段
func (goEngineer GoSoftEngineer) coding(code string) string {
	return "原名" +
		goEngineer.Person.name + "，人称之" +
		goEngineer.name + "," +
		goEngineer.tech + "," +
		(string(goEngineer.experience)) +
		code + ",go"
}
func (jsEngineer JSSoftEngineer) coding(code string) string {
	return "原名" +
		jsEngineer.name + "," +
		jsEngineer.tech + "," +
		(string(jsEngineer.experience)) +
		code + ",js"
}

//func (engineer *GoSoftEngineer) logCode(str string) string {
//	return str
//}
