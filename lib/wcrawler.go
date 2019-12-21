package lib

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var EXCLUDED_MEDIA_EXTENSIONS = []string{
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
}

func (w *WCrawler) Crawler(targetUrl string) *WCrawler {
	resp := Fetcher(targetUrl)

	docs, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	w.ExtractorLink(docs)
	return w
}

func (w *WCrawler) ExtractorLink(docs *goquery.Document) {
	docs.Find("a").Each(func(i int, selection *goquery.Selection) {
		Link, _ := selection.Attr("href")
		w.AllLinks = append(w.AllLinks, Link)
	})

	docs.Find("img").Each(func(i int, selection *goquery.Selection) {
		Link, _ := selection.Attr("src")
		fmt.Println(Link)
		w.AllLinks = append(w.AllLinks, Link)
	})

	docs.Find("iframe").Each(func(i int, selection *goquery.Selection) {
		Link, _ := selection.Attr("src")
		fmt.Println(Link)
		w.AllLinks = append(w.AllLinks, Link)
	})
}

func checkExt(link string) string {
	for _, v := range EXCLUDED_MEDIA_EXTENSIONS {
		if link == v {
			return link
		}
	}
	return ""
}
