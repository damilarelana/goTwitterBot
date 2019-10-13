package main

import (
	"os"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)


// Credentials struct highlighting the fieldnameset required to authenticate each Twitter API client
type Credentials struct{
	consumerKey string
	consumerSecret string
	accessToken string
	accessTokenSecret string
}

// getClient() is the middleware that handles Twitter API sync while returning a useable Twitter Client
func getClient(creds *Credentials)(*twitter.Client, error){
	config := oauth1.NewConfig(creds.consumerKey, creds.consumerSecret)	// stack the consumer key/secret to be used for authentication
	token := oauth1.NewToken(creds.accessToken, creds.accessTokenSecret) // stack the access token/secret to be used for authorisation

	// initiate a Client connection and get a Twitter client instance
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient) 

	// verify login Credentials with Twitter
	verifyParams := %twitter.AccountVerifyParams{
		SkipStatus: twitter.Bool(true),
		IncludeEmail: twitter.Bool(true)
	}

	user, _, err := twitterClient.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Println("Unable to verifiy Twitter credentials")		// print the error encountered
		log.Println(err)		// log the error before returning
		return nil, err
	}

	log.Printf("User's ACCOUNT: \n%+v\n", user) // log the user signing i.e. print to the logger
	return twitterClient, nil
}


func main() {
	fmt.Println("Go Twitter Bot v0.01")

	// extract the environment variables required for Twitter API signing
	creds := Credentials{
		consumerKey := os.Getenv("CONSUMER_KEY")
		consumerSecret := os.Getenv("CONSUMER_SECRET=")
		accessToken := os.Getenv("ACCESS_TOKEN")
		accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	}

	fmt.Printf("%+v\n", creds) // print the credentials in a default format

	twitterClient, err := getClient(&creds) // call the getClient
	if err != nil {
		log.Println("Unable to acquire a Twitter Client")		// log the error before returning
		log.Println(err)
	}


	// sending tweets via the Twitter Client
	tweet, resp, err := twitterClient.Statuses.Update("Golang Twitter Bot sent this ... ", nil)
	if err != nil {
		log.Println("Twitter Bot was unable to post a tweet ... ")
		log.Println(err) // log the error
	}
	log.Printf("%+v\n", resp) // log the response from Twitter
	log.Printf("%+v\n", tweet) // log the tweet that was sent
}
