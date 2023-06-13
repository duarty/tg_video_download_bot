package utils

import "regexp"

func FormatTwitterURL(url string) string {
	regex := regexp.MustCompile(`https[^?]+`)
	return regex.FindString(url)
}