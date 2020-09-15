// Package config provides the configuration.
package config

// Mandatory configuration options.
const (
	SeenFileName = "/Users/Bob/check-youtube/seen-youtube.txt"
	YouTubeKey   = "AIza...Da8c"
)

// GetYouTubeChannelsToCheck returns the channel IDs.
func GetYouTubeChannelsToCheck() []string {
	return []string{
		"UCUORv_qpgmg8N5plVqlYjXg", // Medical Medium
	}
}
