package server

import (
	"fmt"
	"strconv"
)

type Response struct {
	StatusCode int
	Body       []byte
}

var StatusLookUp = map[int]string{
	200: "OK",
	400: "Bad Request",
	404: "Not Found",
	500: "Internal Server Error",
}

func (r Response) Serialize() []byte {
	version := "HTTP/1.1"
	reason := StatusLookUp[r.StatusCode]
	statusLine := version + " " + strconv.Itoa(r.StatusCode) + " " + reason

	header1 := "Content-Type:" + " " + "application/json; charset=utf-8"
	bodyLength := strconv.Itoa(len(r.Body))
	header2 := "Content-Length:" + " " + bodyLength

	responseBytes := fmt.Appendf(nil, "%s\r\n%s\r\n%s\r\n\r\n%s", statusLine, header1, header2, r.Body)

	return responseBytes
}
