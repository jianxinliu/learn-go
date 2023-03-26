package learn_test

import (
	"fmt"
	"strings"
	"testing"
)

type Student struct {
	Name    string
	Age     int
	Address string
}

func (stu Student) ToString() string {
	return fmt.Sprintf("stu[name: %s, age: %d, addr: %s]", stu.Name, stu.Age, stu.Address)
}

func createMap() map[string]Student {
	return map[string]Student{
		"me":  Student{"Jack", 13, "beijing"},
		"her": Student{"rose", 14, "beijing"},
	}
}

func printMap(m map[string]Student) {
	for k, v := range m {
		println(k, v.ToString())
	}
}

func TestMap1(t *testing.T) {
	var stuMap map[string]Student
	stuMap = make(map[string]Student)

	stuMap["me"] = Student{"Jack", 13, "beijing"}
	stuMap["her"] = Student{"rose", 14, "beijing"}

	printMap(stuMap)
}

func TestMap2(t *testing.T) {
	var stuMap = make(map[string]Student)

	stuMap["me"] = Student{"Jack", 13, "beijing"}
	stuMap["her"] = Student{"rose", 14, "beijing"}

	printMap(stuMap)
}

func TestMap3(t *testing.T) {
	var stuMap = map[string]Student{
		"me":  Student{"Jack", 13, "beijing"},
		"her": Student{"rose", 14, "beijing"},
	}

	printMap(stuMap)
}

func TestMap4(t *testing.T) {
	stuMap := createMap()

	delete(stuMap, "me")

	printMap(stuMap)
}

type Writer interface {
	Write(s string)
}

type UpperWriter struct {
	Writer
}

func (UpperWriter) Write(s string) {
	println(strings.ToUpper(s))
}

type LowerWrite struct {
	Writer
}

func (LowerWrite) Write(s string) {
	println(strings.ToLower(s))
}

type JsonWriter struct {
	Writer
}

func (JsonWriter) Write(s string) {
	println(fmt.Sprintf("{%s}\n", s))
}

// Writer types supports
const (
	JSON = iota
	UPPER
	LOWER
)

// write strategy maps
var FnMap = map[int]Writer{
	UPPER: UpperWriter{},
	LOWER: LowerWrite{},
	JSON:  JsonWriter{},
}

// doWrite choose suitable writer for mapType
func doWrite(mapType int, s string) {
	FnMap[mapType].Write(s)
}

func TestStrategyMap(t *testing.T) {
	doWrite(UPPER, "sdkhsdk")
	doWrite(LOWER, "ADSSDjpa123")
	doWrite(JSON, "sdsdgf")
}

func TestMap(t *testing.T) {
	// 声明一个 map 类型的变量
	var stuDb map[string]Student
	// 创建 map
	stuDb = make(map[string]Student)

	// 创建 map 并赋初始化值
	stuDbMap := map[string]Student{
		"me":  Student{"Jack", 13, "beijing"},
		"her": Student{"rose", 14, "beijing"},
	}

	// 也可以直接声明 + 创建
	//var sMap = make(map[string]Student)

	stuDb["me"] = Student{"Jack", 13, "beijing"}
	stuDb["her"] = Student{"rose", 14, "beijing"}

	println("total stu in db: ")
	// 遍历 map
	printMap(stuDbMap)

	me, ok := stuDb["kj"]
	if ok {
		println(me.ToString())
	} else {
		println("wrong key")
	}
}
