package optional

import (
	"testing"

	"github.com/reactivex/rxgo/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOptional(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optional Suite")
}

var _ = Describe("Optional", func() {
	Context("when using non empty optional", func() {
		It("has not nil data", func() {
			testData := "foo"
			some1 := Of(testData)
			got, err := some1.Get()

			Expect(some1.IsEmpty()).NotTo(BeTrue())
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(testData))
		})
		It("has nil data", func() {
			some1 := Of(nil)
			got, err := some1.Get()

			Expect(some1.IsEmpty()).NotTo(BeTrue())
			Expect(err).Should(BeNil())
			Expect(got).Should(BeNil())
		})
	})

	Context("when using Empty Optional", func() {
		It("has empty as the optional", func() {
			empty := Empty()
			got, err := empty.Get()

			Expect(empty.IsEmpty()).Should(BeTrue())
			if err != nil {
				expectedErr := errors.New(errors.NoSuchElementError)
				Expect(err.Error()).Should(Equal(expectedErr.Error()))
			} else {
				Fail("error is not nil")
			}
			Expect(got).Should(BeNil())
		})
	})

})
