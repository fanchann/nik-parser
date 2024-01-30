package nikparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNIKIsValid(t *testing.T) {
	nik := "3321053908775443"
	result, err := ParseNIK(nik)
	assert.Nil(t, err)
	assert.Equal(t, int(1977), result.BornYear)
}

func TestNIKInvalid(t *testing.T) {
	nik := "332105390874241"
	result, err := ParseNIK(nik)
	assert.Error(t, err)
	assert.Nil(t, result)
}
