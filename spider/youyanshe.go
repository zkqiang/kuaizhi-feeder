package spider

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/zkqiang/kuaizhi-feeder/model"
	"net/http"
	"regexp"
	"strconv"
)

func CrawlYouyanshe(page int) (*[]model.Card, error) {
	url := "https://www.yystv.cn/app/home/get_home_feeds_by_page?page=" + strconv.Itoa(page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf(
			"status error: %d %s", resp.StatusCode, resp.Status))
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	code := gjson.Get(doc.Text(), "errorcode").String()
	if code != "20200" {
		return nil, errors.New(fmt.Sprintf(
			"response code error: %s", code))
	}
	var cards []model.Card
	re, _ := regexp.Compile("[\n\t]")
	for _, data := range gjson.Get(doc.Text(), "data").Array() {
		postData := data.Get("data")
		text := postData.Get("text").String()
		if text == "" {
			text = postData.Get("preface").String()
		}
		card := model.Card{
			Images: []string{postData.Get("cover").String()},
			Title:  re.ReplaceAllString(postData.Get("title").String(), ""),
			Text:   re.ReplaceAllString(text, ""),
			Url:    "https://www.yystv.cn/p/" + postData.Get("id").String(),
			Video:  "",
		}
		cards = append(cards, card)
	}
	return &cards, nil
}
