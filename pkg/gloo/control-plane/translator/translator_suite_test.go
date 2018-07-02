package translator

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTranslator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Translator Suite")
}