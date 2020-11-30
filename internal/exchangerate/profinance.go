package exchangerate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://www.profinance.ru/data/cbr-json.php?date=%s"

type Rate struct {
	Quotes map[string]map[string]interface{}
	QuotesNum map[string]map[int]interface{}
	Date string
}

type Profinance struct {

}

func (p *Profinance) GetRate(date string) (float64, error) {
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

	value := rate.Quotes["USD"]["value"].(float64)

	return value, nil
}
