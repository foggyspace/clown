package lib

import (
	"net/http"
)

type Requests struct {
	CyRequest *http.Request
}

func (self *Requests) GetHeaders() http.Header {
	return self.CyRequest.Header
}
