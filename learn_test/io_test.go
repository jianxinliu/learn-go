package learn

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"sort"
	"strconv"
	"testing"
)

func TestReadFile(t *testing.T) {
	lines := ReadFileByLine("./number.log")
	i := lo.Map(lines, func(item string, index int) int {
		atoi, _ := strconv.Atoi(item)
		return atoi
	})
	sort.Ints(i)
	lo.ForEach(i, func(it int, id int) {
		println(it, id)
	})
	WriteToFile("./out.log", i)
}

func ReadFileByLine(fileName string) []string {
	// 开启文件句柄
	file, err := os.Open("./number.log")
	if err != nil {
		panic(err)
	}
	// 关闭文件句柄
	defer file.Close()
	var ret []string
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		ret = append(ret, reader.Text())
	}
	if err := reader.Err(); err != nil {
		panic(err)
	}
	return ret
}

func WriteToFile(dst string, lines []int) {
	// 创建文件句柄，后续可以操作文件
	file, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	// 记得关闭文件句柄
	defer file.Close()
	for _, line := range lines {
		// 向文件输出字符串
		file.WriteString(strconv.Itoa(line) + "\n")
	}
}
