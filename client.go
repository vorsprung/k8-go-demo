package stock

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
)

// Body   head structure
type Body struct {
	Meta   MetaData        `json:"Meta Data"`
	Series TimeSeriesDaily `json:"Time Series (Daily)"`
}

// MetaData   includes symbol
type MetaData struct {
	OneInformation     string `json:"1. Information"`
	TwoSymbol          string `json:"2. Symbol"`
	ThreeLastRefreshed string `json:"3. Last Refreshed"`
	FourOutputSize     string `json:"4. Output Size"`
	FiveTimeZone       string `json:"5. Time Zone"`
}

// OneDay  single day data
type OneDay struct {
	OneOpen               string `json:"1. open"`
	TwoHigh               string `json:"2. high"`
	ThreeLow              string `json:"3. low"`
	FourClose             string `json:"4. close"`
	FiveAdjustedClose     string `json:"5. adjusted close"`
	SixVolume             string `json:"6. volume"`
	SevenDividendAmount   string `json:"7. dividend amount"`
	EightSplitCoefficient string `json:"8. split coefficient"`
}

// TimeSeriesDaily  aggregation
type TimeSeriesDaily map[string]OneDay

// Summary   this is used for output
type Summary struct {
	symbol  string
	data    []float64
	average float64
}

// JSONtoStock  convert incoming text to structs
func JSONtoStock(body []byte) Body {
	var f Body
	err := json.Unmarshal(body, &f)
	if err != nil {
		panic(err)
	}
	return f

}

// Summarize  fill in summary structure
func Summarize(b Body, last int) Summary {
	var rv Summary
	var total float64
	var ordered []string

	rv.symbol = b.Meta.TwoSymbol

	for ymd := range b.Series {
		ordered = append(ordered, ymd)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(ordered)))

	for _, date := range ordered {
		day := b.Series[date]
		closing, err := strconv.ParseFloat(day.FourClose, 64)
		if err != nil {
			panic("float conversion at " + date)
		}
		total = total + closing
		rv.data = append(rv.data, closing)
		last = last - 1
		if last == 0 {
			break
		}
	}
	rv.average = total / float64(len(rv.data))

	return rv
}

// wget   http get and return body
func wget(symbol string) []byte {
	resp := GetReq(symbol)

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		// process body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			return []byte{}
		}
		return body
	}
	fmt.Printf("response code %d\n", resp.StatusCode)

	return []byte{}
}

//GetReq   http GET with some checking
func GetReq(symbol string) *http.Response {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	dud := &http.Response{}
	apikey := os.Getenv("XKEY")
	url := fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s", apikey) +
		"&function=TIME_SERIES_DAILY_ADJUSTED&symbol="

	req, err := http.NewRequest("GET", url+symbol, nil)

	if err != nil {
		fmt.Printf("newrequest: %v\n", err)
	}

	client := &http.Client{}
	if err != nil {
		fmt.Printf("Client Error: %v", err)
		return dud
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Client Do: %v", err)
		return dud
	}
	return resp
}
