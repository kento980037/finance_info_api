package mainIndex

import (
	"testing"
)

func TestScraping(t *testing.T) {
	indexList := []string{"NIKKEI", "TOPIX", "DOW", "SP500", "NASDAQ", "RUSSELL2000", "VIX"}
	for _, index := range indexList {
		dataList := Scraping(index)
		mainIndex := dataList.MainIndex[0]

		code := mainIndex.Code
		if code == "" {
			t.Errorf("code is not get in %v", index)
		}

		name := mainIndex.Name
		if name == "" {
			t.Errorf("name is not get in %v", index)
		}

		price := mainIndex.Price
		if price == "" {
			t.Errorf("price is not get in %v", index)
		}
	}
}
