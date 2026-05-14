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

type customTripper struct {
	Transport http.Transport
	UrlPath   []string
}

func (ct *customTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := ct.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	ct.UrlPath = append(ct.UrlPath, resp.Request.URL.String())
	return resp, nil

}

func Trace(url string) ([]string, error) {

	customTransport := &customTripper{
		Transport: http.Transport{},
	}

	client := &http.Client{
		Transport: customTransport,
	}

	_, err := client.Get(url)
	if err != nil {
		return []string{}, nil
	}

	return customTransport.UrlPath, nil
}
