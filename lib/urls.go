package lib

import (
	"net/url"
	"strings"
)

type CyURL struct {
	CykerURL             url.URL
	Change               bool
	Encoding             string
	AlreadyCalculatedUrl interface{}
}

func HasStartWithSchema(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	} else {
		return "http://" + url
	}
}

func (self *CyURL) GetSchema() string {
	return self.CykerURL.Scheme
}

func (self *CyURL) GetPort() string {
	return self.CykerURL.Port()
}

func (self *CyURL) GetHostName() string {
	return self.CykerURL.Hostname()
}

func (self *CyURL) GetDomain() string {
	return self.CykerURL.Host
}

func (self *CyURL) GetPath() string {
	return self.CykerURL.Path
}

func (self *CyURL) GetRawPath() string {
	return self.CykerURL.RawPath
}

func (self *CyURL) GetFragment() string {
	return self.CykerURL.Fragment
}

func (self *CyURL) GetUrlString() string {
	return self.CykerURL.String()
}
