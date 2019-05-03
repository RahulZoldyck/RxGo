package rxgo

import (

	"github.com/reactivex/rxgo/errors"
)

type Iterator interface {
	cancel()
	Next() (interface{}, error)
}

type iteratorFromChannel struct {
	ch         chan interface{}
	done 		chan interface{}
}

type iteratorFromRange struct {
	current     int
	end         int // Included
	isCancelled bool
}

type iteratorFromSlice struct {
	index       int
	s           []interface{}
	isCancelled bool
}

func (it *iteratorFromChannel) cancel() {
	it.done <- "cancel"
}

func (it *iteratorFromChannel) Next() (interface{}, error) {
	select {
	case <-it.done:
		return nil, errors.New(errors.CancelledIteratorError)
	case next, ok := <-it.ch:
		if ok {
			return next, nil
		}
		return nil, errors.New(errors.EndOfIteratorError)
	}
}

func (it *iteratorFromRange) cancel() {
	it.isCancelled = true
}

func (it *iteratorFromRange) Next() (interface{}, error) {
	if it.isCancelled {
		return nil, errors.New(errors.CancelledIteratorError)
	}
	it.current++
	if it.current <= it.end {
		return it.current, nil
	}
	return nil, errors.New(errors.EndOfIteratorError)
}

func (it *iteratorFromSlice) cancel() {
	it.isCancelled = true
}

func (it *iteratorFromSlice) Next() (interface{}, error) {
	if it.isCancelled {
		return nil, errors.New(errors.CancelledIteratorError)
	}
	it.index++
	if it.index < len(it.s) {
		return it.s[it.index], nil
	}
	return nil, errors.New(errors.EndOfIteratorError)
}

func newIteratorFromChannel(ch chan interface{}) Iterator {
	return &iteratorFromChannel{
		ch:         ch,
		done:        make(chan interface{},1),
	}
}

func newIteratorFromRange(start, end int) Iterator {
	return &iteratorFromRange{
		current:     start,
		end:         end,
		isCancelled: false,
	}
}

func newIteratorFromSlice(s []interface{}) Iterator {
	return &iteratorFromSlice{
		index:       -1,
		s:           s,
		isCancelled: false,
	}
}
