package twitter

import (
	"io"
	"net/http"
	"regexp"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func TwitterDownloader(url string) []byte {


	re := regexp.MustCompile(`status/(\d+)`)
	match := re.FindStringSubmatch(url)

	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet(match[1])
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(tweet.Videos[0].URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	videoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return videoBytes

}
