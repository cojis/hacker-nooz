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

// Story is the story json mapping for HackerNews API
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

// Comment is the comment json mapping for HackerNews API
type Comment struct {
	By     string `json:"by"`
	ID     int    `json:"id"`
	Kids   []int  `json:"kids"`
	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
	Type   string `json:"type"`
}

// TopStories is a type alias for parsing stories
type TopStories = []int

// NoozClient is wrapper around http.Client
type NoozClient struct {
	Client *http.Client
}

// GetComment returns comment of a specific id
func (n *NoozClient) GetComment(id int) Comment {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?", id)
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
	var comment Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		/* TODO: handle */
		panic(err)
	}
	return comment
}

// GetStory returns the HackerNews story of a specific id
func (n *NoozClient) GetStory(id int) Story {
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
	return story
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
		story := n.GetStory(topIds[i])
		fmt.Printf("[%d] %s\n", i, story.Title)
		fmt.Println(story.URL)
		if story.URL != "" {
			fmt.Println()
		}
	}

	var selection int
	fmt.Scanf("%d", &selection)
	story := n.GetStory(topIds[selection])
	for i := 0; i < len(story.Kids); i++ {
		fmt.Println()
		comment := n.GetComment(story.Kids[i])
		fmt.Println(comment.Text)
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
