package nikparser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nikparser "github.com/fanchann/nik-parser"
)

func TestNIKGenWithoutOpts(t *testing.T) {
	niks := nikparser.NIKGen()

	assert.Equal(t, 10, len(niks))
}

func TestNIKGenWithOpts(t *testing.T) {
	nikOpts := &nikparser.NIKOpts{District: "PEKALONGAN BARAT", PostalCode: "51119", Ttl: 100}
	niks := nikparser.NIKGen(nikOpts)

	assert.Equal(t, 100, len(niks))
}

func TestNIKGenDistrictNotFound(t *testing.T) {
	nikOpts := &nikparser.NIKOpts{District: "ATATA 3", PostalCode: "69696", Ttl: 100}
	niks := nikparser.NIKGen(nikOpts)

	assert.Nil(t, niks)
}
