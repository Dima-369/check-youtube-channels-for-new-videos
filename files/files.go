// Package files providers functions around the config.SeenFileName.
package files

import (
	"fmt"
	"io/ioutil"
	"main/config"
	"os"
	"strings"
)

// FetchSeenVideos returns all lines as a string array of the config.SeenFileName file.
func FetchSeenVideos() []string {
	file, err := ioutil.ReadFile(config.SeenFileName)
	if err != nil {
		panic(fmt.Sprintf("ioutil.ReadFile: %v ", err))
	}

	return strings.Split(string(file), "\n")
}

// MarkAsSeen appends the passed video from the channel to the config.SeenFileName file.
func MarkAsSeen(channelName, video string) {
	f, err := os.OpenFile(config.SeenFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(fmt.Sprintf("os.OpenFile: %v ", err))
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(fmt.Sprintf("safeClose: %v", err))
		}
	}()

	if _, err = f.WriteString(createSeenFormat(channelName, video) + "\n"); err != nil {
		panic(fmt.Sprintf("WriteString: %v", err))
	}
}

// HasSeenVideo returns true if the passed video from the channel was already seen.
func HasSeenVideo(seen []string, channelName, video string) bool {
	for _, v := range seen {
		if v == createSeenFormat(channelName, video) {
			return true
		}
	}

	return false
}

func createSeenFormat(channelName, video string) string {
	return fmt.Sprintf("%v|%v", channelName, video)
}
