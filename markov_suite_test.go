package markov_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMarkov(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Markov Suite")
}
