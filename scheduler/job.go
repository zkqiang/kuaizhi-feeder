package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/zkqiang/kuaizhi-feeder/model"
	"net/http"
)

func getJob(token string) *model.Job {
	url := fmt.Sprintf("https://kuaizhi.app/bot/%s/getJob", token)
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("get job token-%s error: %s", token, err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("get job status error: %d %s", res.StatusCode, res.Status)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Errorf("get job token-%s error: %s", token, err)
		return nil
	}
	jobIdData := gjson.Get(doc.Text(), "job.job_id")
	var job *model.Job
	if !jobIdData.Exists() {
		return job
	}
	jobId := jobIdData.String()
	params, _ := gjson.Parse(doc.Text()).Value().(map[string]interface{})
	job = &model.Job{
		JobId:  jobId,
		Params: params,
	}
	return job
}

func postJob(token string, jobId string, cards *[]model.Card) {
	url := fmt.Sprintf("https://kuaizhi.app/bot/%s/pushMessage", token)
	form := &model.FeedForm{
		JobId: jobId,
		Cards: cards,
	}
	formBytes, err := json.Marshal(form)
	if err != nil {
		log.Errorf("post job token-%s id-%s error: %s", token, jobId, err)
		return
	}
	fmt.Print(string(formBytes), url)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(formBytes))
	if err != nil {
		log.Errorf("post job token-%s id-%s error: %s", token, jobId, err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("post job status error: %d %s", res.StatusCode, res.Status)
		return
	}
}
