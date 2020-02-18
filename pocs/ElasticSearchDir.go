package poc

import "fmt"

type ElasticSearchDir struct{}

func (e *ElasticSearchDir) Audit(targetURL string) {
	fmt.Println(e)
}
