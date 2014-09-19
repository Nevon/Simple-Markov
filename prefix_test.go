package markov_test

import (
	. "github.com/nevon/markov"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prefix", func() {
	It("can be joined together", func() {
		p := Prefix{"1", "2"}
		Expect(p.String()).To(Equal("1 2"))
	})

	It("can be shifted", func() {
		p := Prefix{"1", "2"}
		p.Shift("3")
		Expect(p.String()).To(Equal("2 3"))
		p.Shift("4")
		p.Shift("5")
		Expect(p.String()).To(Equal("4 5"))
	})
})
