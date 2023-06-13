package utils

import "regexp"

func IsValidTwitterURL(url string) (bool, string) {
	
	urlFormated := FormatTwitterURL(url)

	twitterUrlPattern := `^https?://twitter\.com/[A-Za-z0-9_]{1,15}/status/[0-9]+(/[?&=A-Za-z0-9_]+)?$`
	r := regexp.MustCompile(twitterUrlPattern)
	return r.MatchString(urlFormated), urlFormated
}
