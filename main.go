package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

// Credentials struct highlighting the fieldnameset required to authenticate each Twitter API client
type Credentials struct {
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
}

// Middleware to help handle twitter API connection negotiation
func getClient(cred *Credentials) (*twitter.Client, error) {

}

func main() {
	consumerKey := os.Getenv("YOUTUBE_KEY")
	consumerSecret := os.Getenv("CHANNEL_ID")
	accessToken := os.Getenv("YOUTUBE_KEY")
	accessTokenSecret := os.Getenv("CHANNEL_ID")
}
