package plugins

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ns3777k/go-shodan/shodan"
)

type ShodanAPI struct{}

func (s *ShodanAPI) getDomain(c *shodan.Client, domain string) {
	info, err := c.GetDomain(context.Background(), domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}

func RunShodan() {
	token := os.Getenv("SHODAN_API_KEY")

	c := shodan.NewClient(nil, token)

	domain := os.Args[1]
	s := ShodanAPI{}
	s.getDomain(c, domain)
}
