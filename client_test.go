// +build http

package stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHttp(t *testing.T) {
	td := wget("ORCL")
	r := JSONtoStock(td)
	assert.Equal(t, r.Meta.TwoSymbol, "ORCL")
}
