package utils

import (
	"net/url"
)

func IsUrlValid(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}
