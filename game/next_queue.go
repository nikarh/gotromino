package game

import "math/rand"

func randomTetromino() Polyomino {
	return tetrominos[rand.Intn(len(tetrominos))]
}

type NextQueue [5]Polyomino

func newNextQueue() *NextQueue {
	queue := new(NextQueue)
	for i := range queue {
		queue[i] = randomTetromino()
	}
	return queue
}

func (n *NextQueue) Take() Polyomino {
	p := n[0]
	for i := 1; i < len(n); i++ {
		n[i-1] = n[i]
	}
	n[len(n)-1] = randomTetromino()

	return p
}
