package search

import (
	"github.com/jgcarvalho/zeca-search-master/rules"
	"github.com/jgcarvalho/zeca/db"
)

type Config struct {
	Title string
	EDA   edaConfig
	Rules rules.Config
	DB    db.Config
	// CA    ca.Config
	Dist distConfig
}

type edaConfig struct {
	Generations int
	Population  int
	Selection   int
	Tournament  int
	OutputProbs string `toml:"output-probabilities"`
	SaveSteps   int    `toml:"save-steps"`
}

type distConfig struct {
	MasterURL string `toml:"master-url"`
	PortA     string `toml:"port-a"`
	PortB     string `toml:"port-b"`
}
