package kabutan

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/xerrors"
)

const (
	URL = "https://kabutan.jp/warning/?mode=3_3&market=0&capitalization=-1&stc=&stm=0&page="

	selectorRoot       = "#main > div.warning_contents > table > tbody > tr"
	selectorCode       = "td:nth-child(1) > a"
	selectorName       = "th"
	selectorMarket     = "td:nth-child(3)"
	selectorPrice      = "td:nth-child(6)"
	selectorDiffPrice  = "td.w61"
	selectorRatioPrice = "td.w50"
)

type Data struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Market     string `json:"market"`
	Price      string `json:"price"`
	DiffPrice  string `json:"diff_price"`
	RatioPrice string `json:"ratio_price"`
}

type DataList struct {
	YearHigh []Data `json:"year_high"`
}

func Scraping() *DataList {
	var yearHigh []Data

	for i := 1; ; i++ {
		// get request
		res, err := http.Get(URL + strconv.Itoa(i))
		if err != nil {
			xerrors.Errorf("http communication failed. :%w", err)
		}
		defer res.Body.Close()

		// reading
		buf, _ := ioutil.ReadAll(res.Body)

		// character code judgement
		det := chardet.NewTextDetector()
		detRslt, _ := det.DetectBest(buf)
		// => EUC-JP

		// character code conversion
		bReader := bytes.NewReader(buf)
		reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

		// HTML parsing
		doc, _ := goquery.NewDocumentFromReader(reader)

		selections := doc.Find(selectorRoot)
		if selections.Length() == 0 {
			break
		}

		selections.Each(func(i int, selection *goquery.Selection) {
			code := selection.Find(selectorCode).Text()
			name := selection.Find(selectorName).Text()
			market := selection.Find(selectorMarket).Text()
			price := selection.Find(selectorPrice).Text()
			diffPrice := selection.Find(selectorDiffPrice).Text()
			ratioPrice := selection.Find(selectorRatioPrice).Text()
			data := Data{Code: code, Name: name, Market: market, Price: price, DiffPrice: diffPrice, RatioPrice: ratioPrice}
			yearHigh = append(yearHigh, data)
		})

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}

	dataList := &DataList{YearHigh: yearHigh}
	return dataList
}
