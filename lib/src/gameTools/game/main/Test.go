package main

import (
	"fmt"
	"reflect"
)

type Peoples interface {
	Show()
}

type A struct{}
type Student struct {	 A * A}
func (stu *Student) Show() {}

func live() Peoples {
	var stu *Student
	return stu
}
const (
	x = "zz"
	y = iota
	z = "zz"
	k
	p = iota
)

type User struct {}

type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
	fmt.Println("MyUser1.m1")
}
func (i User) m2(){
	fmt.Println("User.m2")
}


type Peoples1 interface {
	Speak(string) string
}

type Stduent1 struct{}

//func (stu Stduent) Speak(think string) (talk string) {
func (stu *Stduent1) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {

	//test11()
	//test17()
	//test18()
	//test19()
	//test20()
	//test23()
	//test24();
	//test25()
	//test28()
	test9()
}

func test9() {
	var peo People = &Stduent1{}
	//var peo People = &Stduent{}
	//think := "bitch"
	//fmt.Println(peo.Speak(think))

}

func test28a1() []func() {

	var funs []func()
	for i:=0;i<2 ;i++  {
		funs = append(funs, func() {
			println(&i,i)
		})
	}
	return funs;
}
func test28() {

	funs:=test28a1()
	for _,f:=range funs{
		f()
	}

}

func test25() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}

func test24() {
	type MyInt1 int
	type MyInt2 = int
	//var i int =9
	//var i1 MyInt1 = i 	//基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。 MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转； MyInt2称之为alias，可以直接赋值。
	//var i2 MyInt2 = i
	//fmt.Println(i1,i2)
}

func test23() {
//	for i:=0;i<10 ;i++  {
		//loop:
		//	fmt.println(i)
//	}

	zzb:
		fmt.Println("zzb")
	goto zzb
}

func test20() {
	fmt.Println(x,y,z,k,p)
	//fmt.Println(&p)			//常量不能取址
}

func test19() {
	GetValue := func (m map[int]string, id int) (string, bool) {
		if _, exist := m[id]; exist {
			return "存在数据", true
		}
		//return nil, false
		return "", false
	}

	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}

func test18() {
	Foo := func (x interface{}) {
		if x == nil {
			fmt.Println("empty interface")
			return
		}
		fmt.Println("non-empty interface")
	}

	var x *int = nil
	var y * interface{} = nil

	if y == nil {
		fmt.Println("empty y")
	}

	Foo(x)
}

func test11() {
	s := new(Student)
	fmt.Println(live() , " a: " , s.A)

	if s.A  == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

func test17() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{name: "qq" , age: 11}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{ m: map[string]string{"a": "1"} , age: 11}

	//if sm1 == sm2 {				//结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	//	fmt.Println("sm1 == sm2")
	//}

	if reflect.DeepEqual(sm1 ,sm2) {
		fmt.Println("sm1 == sm2")
	}
}