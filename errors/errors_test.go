package errors

import (
	"testing"

	"github.com/jochasinga/grx/utils/unittest"
	"github.com/stretchr/testify/assert"
)

func TestErrorCodes(t *testing.T) {
	errors := []ErrorCode{
		EndOfIteratorError,
		NilObservableError,
		NilEventStreamError,
		NilObserverError,
		NilSingleError,
		UndefinedError,
	}
	errorEnumTests := unittest.Tables{}
	for i, err := range errors {
		errorEnumTests = append(errorEnumTests, unittest.Table{err, i})
	}

	assert := assert.New(t)
	for _, tt := range errorEnumTests {
		assert.EqualValues(tt.Expected, tt.Actual)
	}
}

func TestBaseErrorImplementError(t *testing.T) {
	baseError := New(EndOfIteratorError)
	assert.Implements(t, (*error)(nil), baseError)
}

func TestBaseErrorWithDefaultMessage(t *testing.T) {
	baseError := New(EndOfIteratorError)
	baseErrorTests := unittest.Tables{
		{baseError.code, ErrorCode(0)},
		{baseError.message, "EndOfIteratorError"},
	}

	assert := assert.New(t)
	for _, tt := range baseErrorTests {
		assert.Equal(tt.Expected, tt.Actual)
	}
}

func TestBaseErrorWithCustomMessage(t *testing.T) {
	baseError := New(NilObservableError, "Observable is set to nil")
	baseErrorTests := unittest.Tables{
		{baseError.code, ErrorCode(1)},
		{baseError.message, "Observable is set to nil"},
	}

	assert := assert.New(t)
	for _, tt := range baseErrorTests {
		assert.Equal(tt.Expected, tt.Actual)
	}
}

func TestErrorMethod(t *testing.T) {
	baseError := New(UndefinedError)
	assert.Equal(t, "4 - UndefinedError", baseError.Error())
}

func TestCodeMethod(t *testing.T) {
	baseError := New(UndefinedError)
	assert.EqualValues(t, 4, baseError.Code())
}
