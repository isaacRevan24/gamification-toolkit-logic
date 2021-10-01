package handlertest

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hello Test suit")
}

var _ = Describe("Hello func", func() {
	Context("Basic test", func() {
		It("Sayed hello isaac", func() {
			Expect(Hello("isaac")).Should(Equal("hello isaac"))
		})
	})
})

func Hello(name string) string {
	return "hello " + name
}
