package solver

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type VerifierMock struct {
	Answers []bool

	currentCall uint
}

func NewVerifierMock() *VerifierMock {
	return &VerifierMock{}
}

func (v *VerifierMock) Verify(token, nonce []byte) (bool, error) {
	if v.currentCall%100 == 0 {
		return true, nil
	}

	return false, nil
}

func TestSolver_Solve(t *testing.T) {

	solver := NewSolver(NewVerifierMock())

	tokens := [][]byte{
		{0, 0, 1, 173, 0, 0, 0, 0, 30, 40, 238, 162, 195, 176, 145, 12},
		{0, 0, 1, 173, 0, 0, 0, 0, 240, 244, 207, 38, 135, 52, 4, 90},
		{0, 0, 1, 173, 0, 0, 0, 0, 115, 2, 237, 22, 160, 129, 174, 7},
	}

	for _, token := range tokens {
		_, err := solver.Solve(context.Background(), token)
		if err != nil {
			t.Error(err)
		}
	}
}
