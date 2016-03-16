package probs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pattern [3]string
type Probability map[string]float64
type ProbRule map[Pattern]Probability

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
		pr[Pattern{ln, c, rn}] = Probability{s1: p1, s2: p2, s3: p3, s4: p4}
	}
	return pr
}
