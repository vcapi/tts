package tts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	googUrl = "https://translate.google.com/translate_tts"
)

// Google tts api
func Google(text, lang string, out io.Writer) error {
	params := make(url.Values)
	params.Add("ie", "UTF-8")
	params.Add("q", text)
	params.Add("tl", lang)
	params.Add("client", "tw-ob")

	addr := fmt.Sprintf("%s?%s", googUrl, params.Encode())
	res, err := request(http.MethodGet, addr, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
