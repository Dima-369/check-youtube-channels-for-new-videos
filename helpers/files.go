package helpers

import (
	"fmt"
	"io/ioutil"
	. "check-youtube/config"
	"os"
	"strings"
)

var Seen []string

func PopulateSeen() {
	file, err := ioutil.ReadFile(SeenFileName)
	if err != nil {
		panic(fmt.Sprintf("ioutil.ReadFile: %v ", err))
	}
	Seen = strings.Split(string(file), "\n")
}

func MarkAsSeen(channelName string, video string) {
	f, err := os.OpenFile(SeenFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("os.OpenFile: %v ", err))
	}
	defer SafeFileClose(f)

	if _, err = f.WriteString(createSeenFormat(channelName, video) + "\n"); err != nil {
		panic(fmt.Sprintf("WriteString: %v", err))
	}
}

func SafeFileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(fmt.Sprintf("SafeFileClose: %v", err))
	}
}

func createSeenFormat(channelName string, video string) string {
	return fmt.Sprintf("%v|%v", channelName, video)
}

func HasSeenVideo(channelName string, video string) bool {
	for _, v := range Seen {
		if v == createSeenFormat(channelName, video) {
			return true
		}
	}
	return false
}
