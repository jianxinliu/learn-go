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
}
func ReadFileByLine(fileName string) []string {
	file, err := os.Open("./number.log")
	if err != nil {
		panic(err)
	}
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
