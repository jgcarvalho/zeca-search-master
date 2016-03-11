package rules

import (
	"testing"

	"github.com/jgcarvalho/zeca-search-master/rules"
)

var conf = rules.Config{
	InitStates:     []string{"#", "A", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "Y"},
	TransStates:    []string{"_", "*", "|", "?"},
	Hydrophobicity: "RosePG"}

func TestCreate(t *testing.T) {
	rule, err := rules.Create(conf)
	if err != nil {
		t.Error(err)
	}
	t.Log(rule)
	t.Log("Code to ###", rule.Code[0][0][0])
}
