package stocktwits

import (
	"fmt"
	"net/http"

	"github.com/khbrendle/stocktwits"
)

type StreamSymbol struct {
	Response stocktwits.Response
	Symbol   stocktwits.Symbol
	Cursor   stocktwits.Cursor
	Messages []stocktwits.Message
}

func FormatUrl(symbol string, params map[string]interface{}) string {
	// would be good to order the map keys to be able to do test on conbinations
	var urlParams string
	var nParams int
	for k, v := range params {
		if nParams == 0 {
			urlParams += "?"
		} else if nParams >= 1 {
			urlParams += "&"
		}
		urlParams += fmt.Sprintf("%s=%v", k, v)
		nParams++
	}
	url := fmt.Sprintf("https://api.stocktwits.com/api/%d/streams/symbol/%s.json%s", stocktwits.APIversion, symbol, urlParams)
	return url
}

func GetStreamSymbol(s string, params map[string]interface{}) (*http.Response, error) {
	var err error
	url := FormatUrl(s, params)
	var resp *http.Response
	// fmt.Println("Getting data")
	if resp, err = http.Get(url); err != nil {
		return resp, err
	}
	return resp, nil
}

func (s StreamSymbol) GetMaxId() int {
	var maxId int
	for _, m := range s.Messages {
		if m.Id > maxId {
			maxId = m.Id
		}
	}
	return maxId
}

func (s StreamSymbol) GetMinId() int {
	var minId int
	for i, m := range s.Messages {
		if i == 0 {
			minId = m.Id
		} else {
			if m.Id < minId {
				minId = m.Id
			}
		}
	}
	return minId
}

// func GetStreamSymbol(s string) (StreamSymbol, error) {
// 	var err error
// 	url := fmt.Sprintf("https://api.stocktwits.com/api/%d/streams/symbol/%s.json", stocktwits.APIversion, s)
// 	var resp *http.Response
// 	fmt.Println("Getting data")
// 	if resp, err = http.Get(url); err != nil {
// 		return StreamSymbol{}, err
// 	}
// 	defer resp.Body.Close()
// 	var body []byte
// 	if body, err = ioutil.ReadAll(resp.Body); err != nil {
// 		return StreamSymbol{}, err
// 	}
// 	var ret StreamSymbol
// 	if err = json.Unmarshal(body, &ret); err != nil {
// 		return StreamSymbol{}, err
// 	}
// 	return ret, nil
// }
