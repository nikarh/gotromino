package game

import (
	"math/rand"
)

type Queue struct{
	Polyominoes [5]Polyomino
	bag []Polyomino
}

func newQueue() *Queue {
	q := new(Queue)
	q.newBag()

	next := q.bag[:5]
	q.bag = q.bag[5:]
	for i := range q.Polyominoes {
		q.Polyominoes[i] = next[i]
	}
	return q
}

func (q *Queue) newBag() {
	l := len(tetrominoes)
	q.bag = make([]Polyomino, l)

	for i, j := range rand.Perm(l) {
		q.bag[i] = tetrominoes[j]
	}
}

func (q *Queue) Take() Polyomino {
	p := q.Polyominoes[0]
	for i := 1; i < len(q.Polyominoes); i++ {
		q.Polyominoes[i-1] = q.Polyominoes[i]
	}

	if len(q.bag) == 0 {
		q.newBag()
	}

	q.Polyominoes[len(q.Polyominoes)-1] = q.bag[0]
	q.bag = q.bag[1:]

	return p
}
