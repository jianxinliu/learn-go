package learn

import (
	"testing"
	"time"
)

// https://coolshell.cn/articles/21146.html

func TestFunctional(t *testing.T) {
	// 填充必要函数，可选函数可以作为可选参数自定义调用对应函数实现
	server := NewServer("localhost", 80, MaxConn(2), Timeout(time.Second*3000))
	println(server)
}

// 使用函数式创建这个对象
type Server struct {
	Addr    string
	Port    int
	MaxConn int
	Timeout time.Duration
	MaxIdle int
}

type Option func(server *Server)

func MaxConn(max int) Option {
	return func(server *Server) {
		server.MaxConn = max
	}
}

func Timeout(time time.Duration) Option {
	return func(server *Server) {
		server.Timeout = time
	}
}

func MaxIdle(max int) Option {
	return func(server *Server) {
		server.MaxIdle = max
	}
}

func NewServer(addr string, port int, ops ...Option) *Server {
	serv := Server{
		Addr: addr,
		Port: port,
		// 以下是默认值
		MaxConn: 100,
		Timeout: time.Second * 30,
		MaxIdle: 10,
	}

	for _, option := range ops {
		option(&serv)
	}
	return &serv
}
