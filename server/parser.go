package server

import (
	"bufio"
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Request struct {
	Method      string
	URL         string
	ProtocolVer string
	Body        []byte
	Headers     map[string]string
	Params      map[string]string
}

var validMethods = []string{"GET", "PUT", "DELETE"}

func Parser(conn io.Reader) (*Request, error) {
	reader := bufio.NewReader(conn)

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("400 bad request- input")
	}

	input = strings.TrimSpace(input)

	parts := strings.SplitN(input, " ", 3)
	if len(parts) != 3 {
		return nil, errors.New("400 bad request- parts")
	}

	method, target, version := parts[0], parts[1], parts[2]

	if !slices.Contains(validMethods, method) {
		return nil, errors.New("400 bad request - invalid method")
	}

	if version != "HTTP/1.1" {
		return nil, errors.New("400 bad request - invalid protocol version")
	}

	headers := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, errors.New("400 bad request - headers")
		}

		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}

		idx := strings.IndexByte(line, ':')
		if idx <= 0 || line[idx-1] == ' ' {
			return nil, errors.New("400 bad request - malformed header")
		}

		//Normalize keys to lowercase for lookup
		key := strings.ToLower(strings.TrimSpace(line[:idx]))
		val := strings.TrimSpace(line[idx+1:])

		headers[key] = val
	}

	req := &Request{
		Method:      method,
		URL:         target,
		ProtocolVer: version,
		Headers:     headers,
	}

	if contentLengthStr, ok := headers["content-length"]; ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil || contentLength < 0 {
			return nil, errors.New("400 bad request - invalid content-length")
		}
		if contentLength > 0 {
			req.Body = make([]byte, contentLength)
			//io.ReadFull guarantees a return of the exact number of bytes in contentLength
			_, err = io.ReadFull(reader, req.Body)
			if err != nil {
				return nil, errors.New("400 bad request - unexpected body EOf")
			}

		}
	}

	return req, nil

}
