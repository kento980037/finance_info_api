package mainIndex

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/xerrors"
)

const (
	URL_ROOT        = "https://www.google.com/finance/quote/"
	URL_NIKKEI      = URL_ROOT + "NI225:INDEXNIKKEI"
	URL_TOPIX       = URL_ROOT + "TOPIX:INDEXTOPIX"
	URL_DOW         = URL_ROOT + ".DJI:INDEXDJX"
	URL_SP500       = URL_ROOT + ".INX:INDEXSP"
	URL_NASDAQ      = URL_ROOT + ".IXIC:INDEXNASDAQ"
	URL_RUSSELL2000 = URL_ROOT + "RUT:INDEXRUSSELL"
	URL_VIX         = URL_ROOT + "VIX:INDEXCBOE"

	selectorPrice = "body > c-wiz > div > div:nth-child(4) > div > div > main > div:nth-child(2) > div:nth-child(1) > c-wiz > div > div:nth-child(1) > div > div:nth-child(1) > div"
)

type Data struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type DataList struct {
	MainIndex []Data `json:"main_index"`
}

func Scraping(code string) *DataList {
	var mainIndex []Data

	indexList := []string{"NIKKEI", "TOPIX", "DOW", "SP500", "NASDAQ", "RUSSELL2000", "VIX"}

	urlMap := make(map[string]string)
	urlMap["NIKKEI"] = URL_NIKKEI
	urlMap["TOPIX"] = URL_TOPIX
	urlMap["DOW"] = URL_DOW
	urlMap["SP500"] = URL_SP500
	urlMap["NASDAQ"] = URL_NASDAQ
	urlMap["RUSSELL2000"] = URL_RUSSELL2000
	urlMap["VIX"] = URL_VIX

	nameMap := make(map[string]string)
	nameMap["NIKKEI"] = "日経平均"
	nameMap["TOPIX"] = "TOPIX"
	nameMap["DOW"] = "ダウ工業平均"
	nameMap["SP500"] = "SP500"
	nameMap["NASDAQ"] = "ナスダック"
	nameMap["RUSSELL2000"] = "ラッセル2000"
	nameMap["VIX"] = "VIX恐怖指数"

	if !isContain(indexList, code) {
		dataList := &DataList{MainIndex: mainIndex}
		return dataList
	}

	// get request
	res, err := http.Get(urlMap[code])
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

	price := doc.Find(selectorPrice).Text()

	data := Data{Code: code, Name: nameMap[code], Price: price}
	mainIndex = append(mainIndex, data)

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

	dataList := &DataList{MainIndex: mainIndex}
	return dataList
}

func isContain(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
