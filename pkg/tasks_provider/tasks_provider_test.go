package tasks_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTasksProvider_Verify(t *testing.T) {
	p := New(0)

	ok, err := p.Verify([]byte{0, 0, 1, 173, 127, 41, 171, 202, 93, 116, 23, 221, 140, 88, 32, 253},
		[]byte{1, 0, 0, 0, 0, 203, 10, 141})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, ok)

	ok, err = p.Verify([]byte{0, 0, 1, 173, 127, 41, 171, 202, 86, 235, 195, 173, 123, 247, 198, 242},
		[]byte{0, 0, 0, 0, 0, 180, 120, 3})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, ok)

	ok, err = p.Verify([]byte{0, 65, 137, 55, 75, 198, 167, 239, 66, 102, 30, 26, 158 ,17, 21, 207},
		[]byte{0, 0, 0, 0, 0, 0, 10, 154})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, ok)
}
