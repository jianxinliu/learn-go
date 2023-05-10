package learn

import (
	"fmt"
	"testing"
)

// slice 源码： /src/runtime/slice.go

func TestSlice(t *testing.T) {
	// 创建一个 slice ,指定长度。
	var sli = make([]string, 5)
	for i := 0; i < len(sli); i++ {
		sli[i] = "aa"
	}
	for i := 0; i < len(sli); i++ {
		println(sli[i])
	}
}

func TestSlice1(t *testing.T) {
	strArr := []string{
		"A", "B", "C",
	}
	// 基于数组创建 slice, 和数组共享底层的数组， slice 的结构包含 len, cap, *data(指向 strArr 底层数组的指针)
	sub := strArr[1:]
	// 会修改底层数组，并且其下标是从 0 重新计算的。这里修改 slice[0] 实际会修改 strArr[1] 的值
	sub[0] = "E"
	for i := 0; i < len(strArr); i++ {
		s := "[empty]"
		if i < len(sub) {
			s = sub[i]
		}
		fmt.Printf("str: %s, slice: %s \n", strArr[i], s)
	}
}

func TestSlice3(t *testing.T) {
	strArr := []string{
		"A", "B", "C",
	}
	sub := strArr[1:]
	println("before append:")
	printArrInfo(&strArr)
	printArrInfo(&sub)

	// 给元素组增加原素，此时会导致原数组扩容，实际的数组变成了另外一个
	strArr = append(strArr, "D")
	println("after append:")
	printArrInfo(&strArr)
	printArrInfo(&sub)
}

// TestAppend append 会往 slice 上增加元素。如果会超过 cap, 则进行扩容，底层数组变成另外一个
// 扩容机制：cap * 2； 详情参考： /src/runtime/slice.go:growslice
// 技巧：如果想实现基于数组创建 slice，但又不想 slice 的操作影响到原数组，可以将 slice 的 len 和 cap 设置成一样的。这样 append 时就会创建新数组
func TestAppend(t *testing.T) {
	slice := make([]int, 0, 15)
	for i := 0; i < 15; i++ {
		slice = append(slice, i)
		fmt.Printf("len:%2d | cap:%2d | v:%2d | addr:%p \n", len(slice), cap(slice), i, &slice)
	}
}

func TestAppendSlice(t *testing.T) {
	arr := []string{"A", "B"}
	slice := arr[:]

	slice = append(slice, "Q")

	slice[0] = "E"
	println(arr[0])
	println(slice[0])
}

func printArrInfo(arr *[]string) {
	fmt.Printf("len: %d, cap: %d, p: %p var: %v \n", len(*arr), cap(*arr), arr, *arr)
}
