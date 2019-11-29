package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// StoryLimit defines the number of stories to show at
// any given time
const StoryLimit = 50

// Story is the json mapping for HackerNews API
type Story struct {
	Creator     string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// TopStories is a type alias for parsing stories
type TopStories = []int

// NoozClient is wrapper around http.Client
type NoozClient struct {
	Client *http.Client
}

// GetStory returns the HackerNews story of a specific id
func (n *NoozClient) GetStory(id int) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-type", `application/json`)

	resp, err := n.Client.Do(req)
	if err != nil {
		/* TODO: handle */
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		/* TODO: handle */
		panic(err)
	}

	// Unmarshal
	var story Story
	err = json.Unmarshal(body, &story)
	if err != nil {
		/* TODO: handle */
		panic(err)
	}
	fmt.Println(story.Title)
	fmt.Println(story.URL)
	fmt.Println()
}

// GetTopStories returns limit number of top stories on HackerNews
func (n *NoozClient) GetTopStories(limit int) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/topstories.json")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-type", `applicaton/json`)

	resp, err := n.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var topIds TopStories
	err = json.Unmarshal(body, &topIds)
	if err != nil {
		panic(err)
	}

	for i := 0; i < limit; i++ {
		n.GetStory(topIds[i])
	}
}

func main() {
	// create client
	n := &NoozClient{
		Client: &http.Client{},
	}
	//	n.GetStory(8863)
	n.GetTopStories(StoryLimit)
}
