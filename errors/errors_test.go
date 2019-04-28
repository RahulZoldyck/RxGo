package errors

import (
	"fmt"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "../matchers"
)

var testErrCodes = []ErrorCode{
	EndOfIteratorError,
	HandlerError,
	ObservableError,
	ObserverError,
	IterableError,
	UndefinedError,
}

func TestErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errors Suite")
}

var _ = Describe("Base Error", func() {
	Context("when creating error", func() {

		It("should implements error interface", func() {
			Expect(New(EndOfIteratorError)).Should(Implements((*error)(nil)))
		})

	})

	Context("when getting error message from Error function", func() {
		It("has no custom message", func() {
			got := New(EndOfIteratorError)
			expectedErrorMsg := fmt.Sprintf("%d - %s", EndOfIteratorError, EndOfIteratorError.String())
			Expect(got.Error()).Should(Equal(expectedErrorMsg))
		})

		It("has 1 custom message", func() {
			msg := "Custom Message"
			got := New(EndOfIteratorError, msg)
			expectedErrorMsg := fmt.Sprintf("%d - %s", EndOfIteratorError, msg)
			Expect(got.Error()).Should(Equal(expectedErrorMsg))
		})

		It("has more than 1 custom message", func() {
			msg1 := "Message 1"
			msg2 := "Message 2"
			got := New(EndOfIteratorError, msg1, msg2)
			expectedErrorMsg := fmt.Sprintf("%d - %s", EndOfIteratorError, msg2)
			Expect(got.Error()).Should(Equal(expectedErrorMsg))
		})

	})

	Context("when getting error code from Code function", func() {
		It("has error code in testErrCodes", func() {
			for _, errCode := range testErrCodes {
				got := New(errCode)
				Expect(got.Code()).Should(Equal(int(errCode)))
			}
		})
	})
})

var _ = Describe("Error Code String", func() {
	It("when the error code is converted into string", func() {
		Expect(EndOfIteratorError.String()).Should(Equal("EndOfIteratorError"))
	})
	It("when error code exceeds the normal set of Errors", func() {
		Expect(CancelledIteratorError.String()).Should(Equal("ErrorCode(10)"))
	})
})
