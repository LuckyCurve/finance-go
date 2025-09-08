package outbound

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CurrencyType string

const (
	USD CurrencyType = "usd"
	CNY CurrencyType = "cny"
	HKD CurrencyType = "hkd"
)

var CollectionCurrencyTypes = []CurrencyType{USD, CNY, HKD}

func GetExchangeRate(time time.Time, currencyType CurrencyType) (map[CurrencyType]float64, error) {
	timeStr := time.Format("2006-01-02") // Go 的日期格式必须写成 "2006-01-02"

	url := fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@%s/v1/currencies/%v.json", timeStr, currencyType)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	res := make(map[CurrencyType]float64)

	for _, c := range CollectionCurrencyTypes {
		res[c] = response.USD[string(c)]
	}

	return res, nil
}

type response struct {
	USD map[string]float64 `json:"usd"`
}
