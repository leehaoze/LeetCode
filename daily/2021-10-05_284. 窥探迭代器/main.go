package main

/*   Below is the interface for Iterator, which is already defined for you.
 *
 */
// type Iterator struct {
// }
//
// func (this *Iterator) hasNext() bool {
// 	// Returns true if the iteration has more elements.
// }
//
// func (this *Iterator) next() int {
// 	// Returns the next element in the iteration.
// }

type PeekingIterator struct {
	iter *Iterator

	oldVal     int
	oldHasNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:       iter,
		oldHasNext: iter.hasNext(),
		oldVal:     iter.next(),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.oldHasNext
}

func (this *PeekingIterator) next() int {
	ret := this.oldVal

	if this.oldHasNext {
		this.oldHasNext = this.iter.hasNext()
		this.oldVal = this.iter.next()
	}

	return ret
}

func (this *PeekingIterator) peek() int {
	return this.oldVal
}
