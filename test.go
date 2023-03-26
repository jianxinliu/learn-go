package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	City string
	*Class
}

type Class struct {
	Name string
}

// 给结构体加方法
func (this *Student) learn() {
	fmt.Printf("stu: %s learning", this.Name)
}

func (this *Class) onClass() {
	fmt.Println(this.Name)
}

func DoTest() {
	fmt.Println("aaa")

	stu := Student{
		Name:  "jack",
		Age:   34,
		City:  "beijing",
		Class: &Class{"class1"},
	}

	stu = Student{
		"rose", 45, "shanghai", &Class{"class1"},
	}

	fmt.Printf("type of *stu %s, &stu: %p \n", reflect.TypeOf(&stu), &stu)

	// & 取一个对象的地址值
	studentGrow(&stu)
	fmt.Println(stu)

	stu.learn()

	// 组合的结构体，可以隐式访问嵌套的结构体内容
	stu.onClass()
	stu.Class.onClass()
}

// * 声明类型为指针类型，表示一个地址值
// 如果对象类型为 Student 而不是 *Student, 则传入的值会被复制，默认是值传递
func studentGrow(s *Student) {
	s.Age++
}
