package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	Error   int
	Message string
	Kind    string
	Data    struct {
		Modhash  string
		Children []struct {
			Kind string
			Data Submission
		}
	}
}

type Submission struct {
	Title         string
	Domain        string
	Url           string
	Author        string
	Score         int
	Selftext_html *string
	Subreddit     string
	Id            string
	Clicked       bool
	Over_18       bool
	Hidden        bool
	Saved         bool
	Is_self       bool
	Edited        interface{} // false or float64
	Thumbnail     string
	Subreddit_id  string
	Downs         int
	Permalink     string
	Name          string
	Created       float64
	Created_utc   float64
	Ups           int
	Num_comments  int
	// num_reports ?
	// distinguished ?
	// banned_by ?
	// media_embed ?
	// author_flair_text ?
	// link_flair_css_class ?
	// author_flair_css_class ?
	// media ?
	// approved_by ?
	// likes ?
	// link_flair_text ?
}

func GetSubmissions(subreddit string) ([]Submission, error) {
	url := fmt.Sprintf("https://reddit.com/r/%s.json", subreddit)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Ilia's Awesome Bot/1.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apiResp apiResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return nil, err
	}
	if apiResp.Error != 0 {
		return nil, errors.New(apiResp.Message)
	}
	ret := make([]Submission, len(apiResp.Data.Children))
	for i, v := range apiResp.Data.Children {
		ret[i] = v.Data
	}

	return ret, nil
}
