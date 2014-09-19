package markov_test

import (
	. "github.com/nevon/markov"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strings"
	"testing"
)

const shortTestString string = "The quick brown fox jumps over the lazy dog"
const longTestString string = "Jag har i alla tider frågat så som vi sagt varför slår mitt hjärta ljublande i takt när med dig just nu jag dansar kind mot kind tusen genom tusen"

var _ = Describe("Chain", func() {
	var (
		singlePrefixChain map[string][]string
		doublePrefixChain map[string][]string
	)

	BeforeEach(func() {
		singlePrefixChain = map[string][]string{
			"":      []string{"The"},
			"the":   []string{"quick", "lazy"},
			"quick": []string{"brown"},
			"brown": []string{"fox"},
			"fox":   []string{"jumps"},
			"jumps": []string{"over"},
			"over":  []string{"the"},
			"lazy":  []string{"dog"},
		}

		doublePrefixChain = map[string][]string{
			"i takt":           []string{"när"},
			"kind tusen":       []string{"genom"},
			"frågat så":        []string{"som"},
			"mitt hjärta":      []string{"ljublande"},
			"ljublande i":      []string{"takt"},
			"dansar kind":      []string{"mot"},
			"slår mitt":        []string{"hjärta"},
			"hjärta ljublande": []string{"i"},
			"takt när":         []string{"med"},
			"kind mot":         []string{"kind"},
			"har i":            []string{"alla"},
			"alla tider":       []string{"frågat"},
			"så som":           []string{"vi"},
			" ":                []string{"Jag"},
			"vi sagt":          []string{"varför"},
			"tusen genom":      []string{"tusen"},
			"mot kind":         []string{"tusen"},
			" jag":             []string{"har"},
			"i alla":           []string{"tider"},
			"tider frågat":     []string{"så"},
			"sagt varför":      []string{"slår"},
			"varför slår":      []string{"mitt"},
			"med dig":          []string{"just"},
			"just nu":          []string{"jag"},
			"jag har":          []string{"i"},
			"dig just":         []string{"nu"},
			"som vi":           []string{"sagt"},
			"när med":          []string{"dig"},
			"nu jag":           []string{"dansar"},
			"jag dansar":       []string{"kind"},
		}

		rand.Seed(1)
	})

	Context("when given a string", func() {
		It("should build a map of prefixes and suffixes", func() {
			c := NewChain(1)
			c.Build(strings.NewReader(shortTestString))

			Expect(c.Chain).To(Equal(singlePrefixChain))
		})

		It("should be able to have different prefix lengths", func() {
			c := NewChain(2)
			c.Build(strings.NewReader(longTestString))
			Expect(c.Chain).To(Equal(doublePrefixChain))
		})
	})

	Context("when given a chain", func() {
		It("should generate text from map", func() {
			c := NewChain(1)
			c.Chain = singlePrefixChain

			Expect(c.Generate(10)).To(Equal("The lazy dog"))
		})

		It("the text should change during subsequent runs", func() {
			c := NewChain(1)
			c.Chain = singlePrefixChain
			Expect(c.Generate(10)).To(Equal("The lazy dog"))
			Expect(c.Generate(10)).To(Equal("The lazy dog"))
			Expect(c.Generate(10)).To(Equal("The quick brown fox jumps over the lazy dog"))
			Expect(c.Generate(10)).To(Equal("The lazy dog"))
		})

		It("should never give more output then asked for", func() {
			c := NewChain(1)
			c.Chain = singlePrefixChain
			Expect(c.Generate(2)).To(Equal("The lazy"))
		})
	})
})

func BenchmarkChainCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewChain(1)
		c.Build(strings.NewReader(longTestString))
	}
}

func BenchmarkSinglePrefixTextGeneration(b *testing.B) {
	c := NewChain(1)
	c.Build(strings.NewReader(longTestString))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Generate(100)
	}
}

func BenchmarkDoublePrefixTextGeneration(b *testing.B) {
	c := NewChain(2)
	c.Build(strings.NewReader(longTestString))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Generate(100)
	}
}
