package stock

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParse   check parsing json is good
func TestParse(t *testing.T) {
	td := helperLoadBytes(t, "query.json")

	r := JSONtoStock(td)

	assert.Equal(t, r.Meta.TwoSymbol, "MSFT")
	assert.Len(t, r.Series, 100)
}

func TestSummary(t *testing.T) {
	td := helperLoadBytes(t, "query.json")
	r := JSONtoStock(td)
	assert.Len(t, r.Series, 100)
	summary := Summarize(r, 3)
	assert.Equal(t, summary.symbol, "MSFT")
	assert.Equal(t, summary.average, 150.08)
	assert.Equal(t, summary.data, []float64{149.31, 149.55, 151.38})

}

func TestWithHttp(t *testing.T) {
	td := wget("ORCL")
	r := JSONtoStock(td)
	assert.Equal(t, r.Meta.TwoSymbol, "ORCL")
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
