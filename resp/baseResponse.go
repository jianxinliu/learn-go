package resp

import "fmt"

type baseResp struct {
	Code    string
	Message string
	Uuid    string
}

func ok() *baseResp {
	return &baseResp{
		"200",
		"",
		"sssssss",
	}
}

func noOk(msg string) *baseResp {
	return &baseResp{
		"500",
		msg,
		"wwwww",
	}
}

type Resp struct {
	Result string
	*baseResp
}

func (this *Resp) ToString() string {
	return fmt.Sprintf("{code: %s, msg: %s, uuid: %s, ret: %s}", this.Code, this.Message, this.Uuid, this.Result)
}

func Success(ret string) *Resp {
	return &Resp{
		ret,
		ok(),
	}
}

func Failed(msg, ret string) *Resp {
	return &Resp{
		ret,
		noOk(msg),
	}
}
