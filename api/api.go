// Package api provides the function to fetch video titles from YouTube's API.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/config"
	"net/http"
)

type activityResponse struct {
	Items []activityItem `json:"items"`
}

type activityItem struct {
	Snippet snippet `json:"snippet"`
}

type snippet struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
}

// FetchResult contains the extracted videos for the YouTube channel.
//
// The ChannelName is also fetched from the YouTube API.
type FetchResult struct {
	ChannelName string
	VideoTitles []string
}

// GetLast5VideoTitlesForChannel delivers its extraction result through the FetchResult channel.
func GetLast5VideoTitlesForChannel(ctx context.Context, channelID string, c chan FetchResult) {
	req, err := http.NewRequestWithContext(ctx, "GET",
		"https://www.googleapis.com/youtube/v3/activities", nil)
	if err != nil {
		panic(fmt.Sprintf("http.NewRequest: %v", err))
	}

	q := req.URL.Query()
	q.Add("key", config.YouTubeKey)
	q.Add("channelId", channelID)
	q.Add("part", "snippet")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("client.Do: %v", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		panic(fmt.Sprintf("ioutil.ReadAll: %v", err))
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status Code is: %v\nResponse body:\n%v", resp.StatusCode, string(body)))
	}

	var response activityResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(fmt.Sprintf("json.Unmarshal: %v ", err))
	}

	videoTitles := make([]string, 0, len(response.Items))

	for _, item := range response.Items {
		videoTitles = append(videoTitles, item.Snippet.Title)
	}

	c <- FetchResult{response.Items[0].Snippet.ChannelTitle, videoTitles}
}
