package matchers

import (
	"github.com/onsi/gomega/types"

	"fmt"
	"reflect"
	"github.com/onsi/gomega/format"
)

func Implements(expected interface{}) types.GomegaMatcher {
	return &ImplementMatcher{
		Expected: expected,
	}
}

type ImplementMatcher struct {
	Expected interface{}
}

func (matcher *ImplementMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	} else if matcher.Expected == nil {
		return false, fmt.Errorf("Refusing to compare type to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	} else if actual == nil {
		return false, nil
	}

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(matcher.Expected).Elem()

	return actualType.Implements(expectedType), nil
}

func (matcher *ImplementMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, fmt.Sprintf("to be implimentation to the type: %T", matcher.Expected))
}

func (matcher *ImplementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, fmt.Sprintf("not to be implimentation to the type: %T", matcher.Expected))
}
