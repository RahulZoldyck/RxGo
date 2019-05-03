package rxgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/reactivex/rxgo/errors"
)

func makeDummyChan() (chan interface{}, Iterator) {
	ch := make(chan interface{}, 1)
	it := newIteratorFromChannel(ch)
	return ch, it
}

func TestIteratorFromChannel(t *testing.T) {
	ch,it := makeDummyChan()

	ch <- 1
	next, err := it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, next)

	ch <- 2
	next, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, next)

	close(ch)
	_, err = it.Next()
	assert.NotNil(t, err)
}

func TestCancellingIteratorFromChannel(t *testing.T) {
	ch := make(chan interface{}, 1)
	it := newIteratorFromChannel(ch)

	ch <- 1
	next, err := it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, next)

	ch <- 2
	next, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, next)

	it.cancel()
	ch <- 3
	next, err = it.Next()
	assert.Nil(t, next)
	assert.Equal(t, errors.New(errors.CancelledIteratorError), err)
}

func TestIteratorFromSlice(t *testing.T) {
	it := newIteratorFromSlice([]interface{}{1, 2, 3})

	next, err := it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, next)

	next, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, next)

	next, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 3, next)

	_, err = it.Next()
	assert.NotNil(t, err)
}

func TestCancellingIteratorFromSlice(t *testing.T) {
	it := newIteratorFromSlice([]interface{}{1, 2, 3})

	next, err := it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, next)

	next, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, next)

	it.cancel()
	next, err = it.Next()
	assert.Nil(t, next)
	assert.Equal(t, errors.New(errors.CancelledIteratorError), err)

}
