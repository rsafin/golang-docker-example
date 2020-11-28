package exchangerate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.profinance.ru/data/cbr-json.php?date=%s"

type Rate struct {
	Quotes interface{}
	QuotesNum interface{}
	Date string
}

type Profinance struct {

}

func (p *Profinance) GetRate(date string) (float32, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rate Rate

	err = json.Unmarshal(body, &rate)
	if err != nil {
		return 0, err
	}

	quotes := rate.Quotes.(map[string]interface{})
	usd := quotes["USD"].(map[string]interface{})
	value := usd["value"]
	fmt.Printf("%+v\n", value)

	return 0, nil
}
