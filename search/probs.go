package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jgcarvalho/zeca-search-master/rules"
)

// type Pattern [3]string
type Probability map[string]float64
type ProbRule map[rules.Pattern]Probability

type Probabilities struct {
	PID        uint32
	Generation int
	Data       ProbRule
}

func ReadProbRule(fn string) ProbRule {
	pr := make(ProbRule)
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("ERROR: reading rule", err)
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var ln, c, rn string
		var s1, s2, s3, s4 string
		var p1, p2, p3, p4 float64

		r := strings.NewReplacer("[", " ", "]", " ", "->", " ", "{", " ", "}", " ", ":", " ", ",", " ")
		fmt.Sscanf(r.Replace(scanner.Text()), "%s %s %s %s %f %s %f %s %f %s %f", &ln, &c, &rn, &s1, &p1, &s2, &p2, &s3, &p3, &s4, &p4)
		pr[rules.Pattern{ln, c, rn}] = Probability{s1: p1, s2: p2, s3: p3, s4: p4}
	}
	return pr
}

func (pk ProbRule) Update(pop []Individual) {
	for pattern, _ := range pk {
		for k, _ := range pk[pattern] {
			pk[pattern][k] = 0.0
		}
	}

	for i := 0; i < len(pop); i++ {
		for pattern, v := range *pop[i].Rule {
			pk[pattern][v] += 1.0 / float64(len(pop))
		}
	}
}
