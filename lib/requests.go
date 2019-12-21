package lib

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"

	"Cyker/utils"
)

func Fetcher(target string) *http.Response {
	request := gorequest.New()
	agent := utils.RandomAgent()
	resp, _, err := request.Get(target).AppendHeader("User-Agent", agent).End()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
