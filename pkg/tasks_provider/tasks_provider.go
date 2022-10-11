package tasks_provider

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type TaskProvider struct {
	Difficulty uint64
}

func New(difficulty uint64) *TaskProvider {
	return &TaskProvider{
		Difficulty: difficulty,
	}
}

func (taskProvider *TaskProvider) NewTask() ([]byte, error) {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], math.MaxUint64/taskProvider.Difficulty)
	_, err := rand.Read(buf[8:])
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return buf, nil
}

func (taskProvider *TaskProvider) Verify(token, nonce []byte) (bool, error) {
	hash := sha256.New()

	_, err := hash.Write(token)
	if err != nil {
		return false, fmt.Errorf("failed to write token: %w", err)
	}
	_, err = hash.Write(nonce)
	if err != nil {
		return false, fmt.Errorf("failed to write nonce: %w", err)
	}

	sum := hash.Sum(nil)
	return bytes.Compare(sum, token) <= 0, nil
}
