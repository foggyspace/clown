package lib

import (
	"golang.org/net/x/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	POST  = "POST"
	GET   = "GET"
	TRACE = "TRACE"
)

type Response *http.Response

type Request *http.Request

type GoCurl struct {
	Url       string
	Method    string
	Header    http.Header
	UserAgent string
	Data      map[string]interface{}
	FormData  http.Values
	QueryData http.Values
	Cookies   []*http.Cookie
	Errors    []error
	BasicAuth struct{ Username, Password string }
	Transport *http.Transport
	Client    *http.Client
	IsClear   bool
}

func (g *GoCurl) NewCurl() {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, _ := cookiejar.New(&cookiejarOptions)

	g := &GoCurl{
		Data:      make(map[string]interface{}),
		Header:    http.Header{},
		FormData:  http.Values{},
		QueryData: http.Values{},
		Client:    &http.Client{Jar: jar},
		Transport: &http.Transport{},
		BasicAuth: struct{ Username, Password string }{},
		Cookies:   make([]*http.Cookie, 0),
		Errors:    nil,
	}

	g.Transport.DisableKeepAlives = true
	return g
}

func (g *GoCurl) ClearGoCurl() {
	if g.IsClear {
		return
	}
	g.Url = ""
	g.Method = ""
	g.Header = http.Header{}
	g.FormData = http.Values{}
	g.QueryData = http.Values{}
	g.Data = make(map[string]interface{})
	g.Errors = nil
	g.Cookies = make([]*http.Cookie, 0)
}

func (g *GoCurl) Get(rawUrl string) *GoCurl {
	g.ClearGoCurl()
	g.Method = GET
	g.Url = rawUrl
	g.Error = nil
	return g
}

func (g *GoCurl) Post(rawUrl string) *GoCurl {
	g.ClearGoCurl()
	g.Method = POST
	g.Url = rawUrl
	g.Error = nil
	return g
}

func (g *GoCurl) Trace(rawUrl string) *GoCurl {
	g.ClearGoCurl()
	g.Method = TRACE
	g.Url = rawUrl
	g.Error = nil
	return g
}

func (g *GoCurl) SetHeader(key, value string) *GoCurl {
	g.Header.Set(key, value)
	return g
}

func (g *GoCurl) AddHeader(key, value string) *GoCurl {
	g.Header.Add(key, value)
	return g
}

func (g *GoCurl) SetProxy(proxyUrl string) *GoCurl {
	proxiesUrl, err := url.Parse(proxyUrl)
	if err != nil {
		g.Error = append(s.Error, err)
	} else if proxyUrl == "" {
		g.Transport.Proxy = nil
	} else {
		g.Transport.Proxy = http.ProxyURL(proxiesUrl)
	}
	return g
}

func (g *GoCurl) getResponseBytes() (Response, []byte, error) {}
