package main

import (
	"context"
	"fmt"
	"main/api"
	"main/config"
	"main/files"
)

func hasNewVideos(seen []string, result api.FetchResult) bool {
	for _, title := range result.VideoTitles {
		if !files.HasSeenVideo(seen, result.ChannelName, title) {
			return true
		}
	}

	return false
}

func main() {
	ctx := context.Background()
	youTubeChannelIDs := config.GetYouTubeChannelsToCheck()
	seen := files.FetchSeenVideos()

	channels := make([]chan api.FetchResult, 0, 64)

	for _, youTubeChannel := range youTubeChannelIDs {
		c := make(chan api.FetchResult)

		channels = append(channels, c)

		go api.GetLast5VideoTitlesForChannel(ctx, youTubeChannel, c)
	}

	for i, c := range channels {
		result := <-c

		if hasNewVideos(seen, result) {
			channelID := youTubeChannelIDs[i]

			fmt.Printf("%v   https://www.youtube.com/channel/%v\n",
				result.ChannelName, channelID)

			for _, title := range result.VideoTitles {
				p := fmt.Sprintf("  - %v", title)

				if !files.HasSeenVideo(seen, result.ChannelName, title) {
					files.MarkAsSeen(result.ChannelName, title)
					fmt.Println(p)
				}
			}

			println("")
		}
	}
}
