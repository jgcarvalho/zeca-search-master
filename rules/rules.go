package rules

type Config struct {
	Input  string `toml:"input"`
	Output string `toml:"output"`
}

type Pattern [3]string
type Rule map[Pattern]string
