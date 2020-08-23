package main

import (
	"fmt"
	. "check-youtube/config"
	. "check-youtube/helpers"
)

func HasNewVideos(result FetchResult) bool {
	for _, title := range result.VideoTitles {
		if !HasSeenVideo(result.ChannelName, title) {
			return true
		}
	}
	return false
}

func main() {
	PopulateSeen()

	var channels []chan FetchResult
	for _, youTubeChannel := range YouTubeChannelIds {
		c := make(chan FetchResult)
		channels = append(channels, c)
		go GetLast5VideoTitlesForChannel(youTubeChannel, c)
	}

	for i, c := range channels {
		result := <-c

		if HasNewVideos(result) {
			channelId := YouTubeChannelIds[i]

			fmt.Printf("%v   https://www.youtube.com/channel/%v\n",
				result.ChannelName, channelId)

			for _, title := range result.VideoTitles {
				p := fmt.Sprintf("  - %v", title)
				if !HasSeenVideo(result.ChannelName, title) {
					MarkAsSeen(result.ChannelName, title)
					fmt.Println(p)
				}
			}
			println("")
		}
	}
}
