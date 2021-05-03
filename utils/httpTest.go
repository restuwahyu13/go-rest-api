package util

import (
	"bytes"
	"encoding/binary"
	"net/http"
	"net/http/httptest"
)

func HttpTestRequest(method, url string, payload []byte) (*httptest.ResponseRecorder, *http.Request, error) {

	request := make(chan *http.Request, 1)
	errors := make(chan error, 1)

	if binary.Size(payload) > 0 {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		request <- req
		errors <- err
	} else {
		req, err := http.NewRequest(method, url, nil)
		request <- req
		errors <- err
	}

	rr := httptest.NewRecorder()

	return rr, <-request, <-errors
}
