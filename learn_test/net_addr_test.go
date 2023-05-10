package learn

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestNetAddr(t *testing.T) {
	SubNet("B", 10)

	SubNet("A", 5)

	SubNet("C", 5)
	SubNet("C", 4)
}

var NetMap = map[string]int{
	"A": int(math.Pow(2, 1*8)),
	"B": int(math.Pow(2, 2*8)),
	"C": int(math.Pow(2, 3*8)),
}

func toBin(i int) string {
	return strconv.FormatInt(int64(i), 2)
}

func SubNet(netType string, subCnt int) {
	hostId := int(32 - math.Log2(float64(NetMap[netType])))

	totalNetCnt := 0
	for i := 0; i < subCnt; i++ {
		hostCnt := hostId - i
		currentSubNetCnt := int(math.Pow(2, float64(i+hostCnt)))
		fmt.Printf("当前位数最大子网ID: %s, 当前位数最大主机ID: %s, 当前位数最大主机数: %d\n",
			toBin(int(math.Pow(2, float64(i)))),
			toBin(hostCnt),
			currentSubNetCnt)
		totalNetCnt += currentSubNetCnt
	}
	fmt.Printf("%s 类网络， %d 位子网, 总主机数: %d, 比不分子网主机数多出： %f%% \n\n", netType, subCnt, totalNetCnt,
		(math.Pow(2, 32)/float64(totalNetCnt))*100.0)
}
