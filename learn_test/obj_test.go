package learn

import "testing"

type Student struct {
	Addr string
}

func (s Student) ToString() string {
	return s.Addr
}

// 相当于静态方法
func (Student) A() string {
	return ""
}

func (s Student) B() string {
	s.Addr = "asa"
	return s.Addr
}

// 需要改变对象的值，就需要传引用
func (s *Student) C() string {
	s.Addr = "asasd"
	return s.Addr
}

func TestObject(t *testing.T) {
	stu := Student{
		Addr: "sd",
	}
	stu.A()
	println(stu.ToString())
	stu.B()
	println(stu.ToString())
	stu.C()
	println(stu.ToString())
}
