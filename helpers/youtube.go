package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "check-youtube/config"
	"net/http"
)

type ActivityResponse struct {
	Items []ActivityItem `json:"items"`
}

type ActivityItem struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
}

type FetchResult struct {
	ChannelName string
	VideoTitles []string
}

func GetLast5VideoTitlesForChannel(channelId string, c chan FetchResult) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/activities", nil)
	if err != nil {
		panic(fmt.Sprintf("http.NewRequest: %v", err))
	}

	q := req.URL.Query()
	q.Add("key", YouTubeKey)
	q.Add("channelId", channelId)
	q.Add("part", "snippet")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("client.Do: %v", err))
	}
	defer SafeBodyClose(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("ioutil.ReadAll: %v", err))
	}

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Status Code is: %v\nResponse body:\n%v", resp.StatusCode, string(body)))
	}

	var response ActivityResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(fmt.Sprintf("json.Unmarshal: %v ", err))
	}

	var videoTitles []string
	for _, item := range response.Items {
		videoTitles = append(videoTitles, item.Snippet.Title)
	}

	c <- FetchResult{response.Items[0].Snippet.ChannelTitle, videoTitles}
}

func SafeBodyClose(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("SafeBodyClose: %v", err))
	}
}
