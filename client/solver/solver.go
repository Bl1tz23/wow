package solver

import (
	"context"
	"encoding/binary"
	"errors"
)

type Solver struct {
	v Verifier
}

func NewSolver(v Verifier) *Solver {
	return &Solver{
		v,
	}
}

type Verifier interface {
	Verify(token, nonce []byte) (bool, error)
}

func (s *Solver) Solve(ctx context.Context, token []byte) ([]byte, error) {
	solutionChan := make(chan []byte)

	go s.solve(token, solutionChan)

	select {
	case <-ctx.Done():
		return nil, errors.New("deadline exceeded")
	case solution := <-solutionChan:
		return solution, nil
	}
}

// bruteforce to find such hash(nonce + token) <= token
func (s *Solver) solve(token []byte, solution chan []byte) {
	nonce := make([]byte, 8)
	for i := uint64(0); ; i++ {
		binary.BigEndian.PutUint64(nonce, i)
		if ok, err := s.v.Verify(token, nonce); ok && err == nil {
			solution <- nonce
			return
		}
	}
}
