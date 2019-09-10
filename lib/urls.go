package lib

import (
	"net/url"
	"strings"
)

type CyURL struct {
	Change               bool
	Encoding             string
	AlreadyCalculatedUrl interface{}
}

func ParserUrl(rawUrl string) (*url.URL, error) {
	if strings.HasPrefix(rawUrl, "http://") || strings.HasPrefix(rawUrl, "https://") {
		cyurl, _ := url.Parse(rawUrl)
		return cyurl, nil
	} else {
		raw_url := "http://" + rawUrl
		cyurl, _ := url.Parse(raw_url)
		return cyurl, nil
	}
}

func (self *CyURL) GetSchema(u *url.URL) string {
	return u.Scheme
}

func (self *CyURL) GetPort(u *url.URL) string {
	return u.Port()
}

func (self *CyURL) GetHostName(u *url.URL) string {
	return u.Hostname()
}

func (self *CyURL) GetDomain(u *url.URL) string {
	return u.Host
}

func (self *CyURL) GetPath(u *url.URL) string {
	return u.Path
}

func (self *CyURL) GetRawPath(u *url.URL) string {
	return u.RawPath
}

func (self *CyURL) GetFragment(u *url.URL) string {
	return u.Fragment
}

func (self *CyURL) GetUrlString(u *url.URL) string {
	return u.String()
}

func (self *CyURL) GetQuery(u *url.URL) url.Values {
	return u.Query()
}

func (self *CyURL) GetRawQuery(u *url.URL) string {
	return u.RawQuery
}

func (self *CyURL) GetRequestURI(u *url.URL) string {
	return u.RequestURI()
}
