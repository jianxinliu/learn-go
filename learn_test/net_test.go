package learn

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"testing"
)

func TestNet(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	retStr := "GET /v1/hello HTTP/1.1\r\nHost: 127.0.0.1:8080\r\n\r\n"
	//retStr = "HEAD /v1/hello HTTP/1.1 \r\n\r\n"

	_, err = conn.Write([]byte(retStr))
	if err != nil {
		panic(err)
	}

	ret, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	println(string(ret))
}

func TestHttp(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/v1/hello")
	if err != nil {
		panic(err)
	}
	ret, err := io.ReadAll(resp.Body)
	println(string(ret))
}

func readAll(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result.Write(buf[0:n])
	}
	return result.Bytes(), nil
}
