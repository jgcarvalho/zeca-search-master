package search

import (
	"testing"

	"github.com/jgcarvalho/zeca-search-master/search"
)

func TestReadProbRule(t *testing.T) {
	fn := "/home/jgcarvalho/gocode/src/github.com/jgcarvalho/zeca-create-rule/rose.rule"
	pr := search.ReadProbRule(fn)
	for k, v := range pr {
		t.Log(k, v)
	}
}
