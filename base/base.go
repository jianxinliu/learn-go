package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var scores [10]int

func Init() {
	scores[0] = 100

	fmt.Println(scores)

	scores1 := []int{1, 2, 3}

	fmt.Println(scores1)

	fmt.Println(len(scores))

	for _, score := range scores {
		fmt.Println(score)
	}

	ages := make([]int, 2, 10)
	fmt.Println(ages)
}

type BaseRet struct {
	Flag    bool   `json:"flag"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
	Uid     string `json:"uid"`
}

func Success(c *gin.Context, ret any) {
	c.JSON(200, BaseRet{
		true, http.StatusOK,
		"", ret, "",
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, BaseRet{
		false, code,
		msg, nil, "",
	})
}
