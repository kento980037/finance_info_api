package kabutan

import (
	"testing"
)

func TestScraping(t *testing.T) {
	dataList := Scraping()
	for _, year_high := range dataList.YearHigh {
		code := year_high.Code
		if code == "" {
			t.Errorf("code is not get")
		}

		name := year_high.Name
		if name == "" {
			t.Errorf("name is not get")
		}

		market := year_high.Market
		if market == "" {
			t.Errorf("market is not get")
		}

		price := year_high.Price
		if price == "" {
			t.Errorf("price is not get")
		}

		diff_price := year_high.DiffPrice
		if diff_price == "" {
			t.Errorf("diff_price is not get")
		}

		ratio_price := year_high.RatioPrice
		if ratio_price == "" {
			t.Errorf("ratio_price is not get")
		}
	}
}
