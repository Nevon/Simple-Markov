package markov

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
)

type Chain struct {
	Chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Build(r io.Reader) {
	br := bufio.NewScanner(r)
	p := make(Prefix, c.prefixLen)
	br.Split(bufio.ScanWords)
	for br.Scan() {
		s := br.Text()
		key := strings.ToLower(p.String())
		c.Chain[key] = append(c.Chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.Chain[strings.ToLower(p.String())]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}
