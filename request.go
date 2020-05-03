package tts

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:74.0) Gecko/20100101 Firefox/74.0"
	httpTimeout = 5 * time.Second
)

func request(ctx context.Context, method, addr string, body io.Reader) (*http.Response, error) {
	up, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, addr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Host", up.Host)
	client := http.Client{
		Timeout: httpTimeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed: %s", res.Status)
	}
	return res, nil
}
