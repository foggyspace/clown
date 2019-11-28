package lib

import (
	"github.com/PuerkitoBio/goquery"
)

const EXCLUDED_MEDIA_EXTENSIONS = []string{
	".7z", ".aac", ".aiff", ".au", ".avi", ".bin", ".bmp", ".cab", ".dll", ".dmp", ".ear", ".exe", ".flv", ".gif",
	".gz", ".image", ".iso", ".jar", ".jpeg", ".jpg", ".mkv", ".mov", ".mp3", ".mp4", ".mpeg", ".mpg", ".pdf", ".png",
	".ps", ".rar", ".scm", ".so", ".tar", ".tif", ".war", ".wav", ".wmv", ".zip",
}

type WCrawler struct {
	AllLinks []string
	OkLinks  []string
	Forms    []string
	Scheme   string
	Netloc   string
	Response []byte
}

func (w *WCrawler) Crawler(targetUrl string) []string {}

func (w *WCrawler) ExtractorLink() {}

func (w *WCrawler) CheckExt(link string) string {}

func (w *WCrawler) CheckUrl(link string) string {}
