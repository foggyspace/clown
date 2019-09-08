package lib

var (
	DEFAULTENCODING = "utf-8"
	CR              = "\r"
	LF              = "\n"
	CRLF            = "\r\n"
	SP              = " "
)

type CyResponse struct {
	Code    int64
	Headers map[string]interface{}
	Charset string
}

func (self *CyResponse) GetCode() int64 {
	return self.Code
}

func (self *CyResponse) SetCode(code int64) {
	self.Code = code
}

func (self *CyResponse) GetCharset() string {
	return self.Charset
}

func (self *CyResponse) SetCharset(chasert string) {
	self.Charset = chasert
}
