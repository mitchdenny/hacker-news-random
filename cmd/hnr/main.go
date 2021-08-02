package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/mitchdenny/hacker-news-random/hnapi"
)

func printHackerNewsItem(item hnapi.HackerNewsItem) {
	if item.ParentId == 0 {
		itemTemplate := "%d (%s): %s\n"
		fmt.Printf(itemTemplate, item.Id, item.User, item.Url)
	} else {
		itemTemplate := "\u2515 %d (%s): \n%s\n\n"
		fmt.Printf(itemTemplate, item.Id, item.User, item.Text)
	}
}

func getParentHackerNewsItems(item hnapi.HackerNewsItem) ([]hnapi.HackerNewsItem, error) {
	var parentItems []hnapi.HackerNewsItem
	parentItem := item
	parentItems = append(parentItems, parentItem)

	for parentItem.ParentId != 0 {
		var err error
		parentItem, err = hnapi.GetItem(parentItem.ParentId)
		if err != nil {
			return []hnapi.HackerNewsItem{}, err
		}

		parentItems = append(parentItems, parentItem)
	}

	sort.Slice(parentItems[:], func(i, j int) bool {
		return parentItems[i].Id < parentItems[j].Id
	})

	return parentItems, nil
}

func getRandomKidHackerNewsItems(item hnapi.HackerNewsItem) ([]hnapi.HackerNewsItem, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomKidHackerNewsItems := []hnapi.HackerNewsItem{}

	previousHackerNewsItem := item
	for len(previousHackerNewsItem.Kids) > 0 {
		randomKidIndex := r.Intn(len(previousHackerNewsItem.Kids))
		selectedKidHackerNewsItemIndex := previousHackerNewsItem.Kids[randomKidIndex]
		selectedKidHackerNewsItem, err := hnapi.GetItem(selectedKidHackerNewsItemIndex)
		if err != nil {
			return []hnapi.HackerNewsItem{}, err
		}

		randomKidHackerNewsItems = append(randomKidHackerNewsItems, selectedKidHackerNewsItem)
		previousHackerNewsItem = selectedKidHackerNewsItem
	}

	return randomKidHackerNewsItems, nil
}

func main() {
	maxItemId, err := hnapi.GetMaxItemId()
	if err != nil {
		log.Fatal(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomItemId := r.Intn(maxItemId)

	item, err := hnapi.GetItem(randomItemId)
	if err != nil {
		log.Fatal(err)
	}

	parentHackerNewsItems, err := getParentHackerNewsItems(item)
	if err != nil {
		log.Fatal(err)
	}

	randomKidHackerNewsItems, err := getRandomKidHackerNewsItems(item)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range parentHackerNewsItems {
		printHackerNewsItem(v)
	}

	for _, v := range randomKidHackerNewsItems {
		printHackerNewsItem(v)
	}
}
