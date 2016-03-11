package rules

import (
	"fmt"
	"math/rand"
	"time"
)

type Config struct {
	Input          string   `toml:"input"`
	Output         string   `toml:"output"`
	InitStates     []string `toml:"initial-states"`
	TransStates    []string `toml:"transition-states"`
	Hydrophobicity string   `toml:"hydrophobicity"`
}

// Rule represents a CA rule
type Rule struct {
	Code  [][][]byte
	Fixed [][][]bool
	Prm   Params
}

// Params represents the parameters of the CA rule, with information about start
// and transition rules, if there is a "wild card" in the transition states, and
// the neighborhood size R
type Params struct {
	StrStates           []string
	StrTransitionStates []string
	States              []byte
	TransitionStates    []byte
	Hydrophobicity      string
	// R                uint8
}

func getHydroPrm(conf Config) (hStates []string) {
	scale := conf.Hydrophobicity
	if scale == "None" {
		hStates = make([]string, len(conf.TransStates)-1) //exclui o estado ?
		for i := 0; i < len(hStates); i++ {
			hStates[i] = conf.TransStates[i]
		}
	} else if scale == "RosePG" {
		hStates = make([]string, (len(conf.TransStates)-1)*4) //exlui o estado ?
		for i := 0; i < (len(conf.TransStates) - 1); i++ {
			hStates[i] = conf.TransStates[i] + "n"
			hStates[i+3] = conf.TransStates[i] + "p"
			hStates[i+6] = conf.TransStates[i] + "P"
			hStates[i+9] = conf.TransStates[i] + "G"
		}
	} else if scale == "RosePGCharged" {
		hStates = make([]string, (len(conf.TransStates)-1)*6) //exclui o estado ?
		for i := 0; i < (len(conf.TransStates) - 1); i++ {
			hStates[i] = conf.TransStates[i] + "n"
			hStates[i+3] = conf.TransStates[i] + "p"
			hStates[i+6] = conf.TransStates[i] + "P"
			hStates[i+9] = conf.TransStates[i] + "G"
			hStates[i+12] = conf.TransStates[i] + "+"
			hStates[i+15] = conf.TransStates[i] + "-"
		}
	} else {
		panic("Hydrophobicity scale not found!!!")
	}
	return
}

// Create a new rule given the start states, transition states, if there is a joker
// (which must be the last element in the transition states) and neighborhood r.
func Create(conf Config) (*Rule, error) {
	rand.Seed(time.Now().UnixNano())
	var ru Rule

	hydroStates := getHydroPrm(conf)
	fmt.Println(hydroStates)
	ru.Prm.States = make([]byte, len(conf.InitStates)+len(hydroStates))
	ru.Prm.StrStates = make([]string, len(conf.InitStates)+len(hydroStates))
	for i := 0; i < len(ru.Prm.States); i++ {
		ru.Prm.States[i] = byte(i)
		if i < len(conf.InitStates) {
			ru.Prm.StrStates[i] = conf.InitStates[i]
		} else {
			ru.Prm.StrStates[i] = hydroStates[i-len(conf.InitStates)]
		}
	}
	ru.Prm.TransitionStates = make([]byte, len(conf.TransStates))
	ru.Prm.StrTransitionStates = make([]string, len(conf.TransStates))
	for i := 0; i < len(conf.TransStates); i++ {
		ru.Prm.TransitionStates[i] = byte(len(ru.Prm.States) + i)
		ru.Prm.StrTransitionStates[i] = conf.TransStates[i]
	}

	sst := ru.Prm.States
	tst := ru.Prm.TransitionStates

	ru.Code = make([][][]byte, len(sst))
	ru.Fixed = make([][][]bool, len(sst))
	for ln := range sst {
		ru.Code[ln] = make([][]byte, len(sst))
		ru.Fixed[ln] = make([][]bool, len(sst))
		for c := range sst {
			ru.Code[ln][c] = make([]byte, len(sst))
			ru.Fixed[ln][c] = make([]bool, len(sst))
			for rn := range sst {
				ru.Code[ln][c][rn] = tst[rand.Intn(len(tst))]
				// c == 0 usually means # that represents the n and c terminal. it's not a residue
				// but it's essencial to CA representation and must be fixed (never change)
				if c == 0 {
					ru.Code[ln][c][rn] = 0
					ru.Fixed[ln][c][rn] = true
				} else {
					ru.Fixed[ln][c][rn] = false
				}
			}
		}
	}
	return &ru, nil
}

func (r *Rule) String() string {
	codes := append(r.Prm.StrStates, r.Prm.StrTransitionStates...)
	var toprint string
	for c := 0; c < len(r.Code); c++ {
		for ln := 0; ln < len(r.Code); ln++ {
			for rn := 0; rn < len(r.Code); rn++ {
				toprint += fmt.Sprintf("[%s][%s][%s] -> [%s]\n", codes[ln], codes[c], codes[rn], codes[r.Code[ln][c][rn]])
				// toprint += fmt.Sprintf("[%d][%d][%d] -> [%d]\n", ln, c, rn, r.Code[ln][c][rn])
			}
		}
	}
	return toprint
}
