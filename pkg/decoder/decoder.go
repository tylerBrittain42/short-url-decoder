package decoder

import (
	"net/http"
)

func FinalDestination(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Request.URL.String(), nil
}

func Trace(url string) ([]string, error) {
	return []string{url}, nil
}
