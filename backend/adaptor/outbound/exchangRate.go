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

var collectionCurrencyTypes = []CurrencyType{CNY, HKD}

func GetExchangeRate(time time.Time) (map[CurrencyType]float64, error) {
	timeStr := time.Format("2006-01-02") // Go 的日期格式必须写成 "2006-01-02"

	// 拼接 URL
	url := fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@%s/v1/currencies/usd.json", timeStr)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
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

	for _, c := range collectionCurrencyTypes {
		res[c] = response.USD[string(c)]
	}

	return res, nil
}

type response struct {
	USD map[string]float64 `json:"usd"`
}
