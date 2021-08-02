package hnapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HackerNewsItem struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent"`
	User     string `json:"by"`
	Text     string `json:"text"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	Kids     []int  `json:"kids"`
}

func GetMaxItemId() (int, error) {
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json?print=pretty")
	if err != nil {
		return -1, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	responseString := string(responseData)

	trimmedResponseString := strings.Trim(responseString, "\n")

	responseInt, err := strconv.Atoi(trimmedResponseString)
	if err != nil {
		return -1, err
	}

	return responseInt, nil
}

func GetItem(itemId int) (HackerNewsItem, error) {
	requestUrlTemplate := "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
	requestUrl := fmt.Sprintf(requestUrlTemplate, itemId)
	response, err := http.Get(requestUrl)
	if err != nil {
		return HackerNewsItem{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HackerNewsItem{}, err
	}

	var item HackerNewsItem
	if err := json.Unmarshal(responseData, &item); err != nil {
		return HackerNewsItem{}, err
	}

	return item, nil
}
