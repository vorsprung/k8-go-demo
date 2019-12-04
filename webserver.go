package stock

import (
	"fmt"
	"net/http"
	"strconv"
)

// SummaryResponse   application response for webserver
func SummaryResponse(r *http.Request) string {
	q := r.URL.Query()
	symbol := q.Get("symbol")
	ndays, err := strconv.ParseInt(q.Get("ndays"), 10, 32)
	if err != nil {
		ndays = 3
	}
	td := wget(symbol)
	bodystruct := JSONtoStock(td)
	summary := Summarize(bodystruct, int(ndays))

	return fmt.Sprintf("%s data=%v, average=%f",
		symbol,
		summary.data,
		summary.average)
}
