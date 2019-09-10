package common

import "regexp"

const (
	URLHEADER = "location"
)

var (
	URLTAGS   = [...]string{"a", "img", "link", "script", "iframe", "object", "embed", "area", "frame", "applet", "input", "base", "div", "layer", "form"}
	URLATTRS  = [...]string{"href", "src", "data", "action"}
	SAFECHARS = [...]string{"\x00", "%00"}
	URLRE     = regexp.MustCompile(`((http|https)://([\w:@\-\./]*?)[^ \n\r\t"\'<>)\s]*)`)
)

type HtmlParser struct {
	Encoding string
}

func (self *HtmlParser) GetUrlFromResponseHeader() string {}

func (self *HtmlParser) GetUrlFromTag() string {}

func (self *HtmlParser) ExtractorGoodUrl() string {}
