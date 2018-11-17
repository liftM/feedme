package caviar

import (
	"net/http"
	"net/url"
	"strings"
)

func caviarURL(url string) string {
	return URL + "/" + strings.TrimPrefix(url, "/")
}

// PostForm behaves *http.Client.PostForm while automatically setting the URL
// prefix and CSRF token.
func (s *Session) PostForm(url string, data url.Values) (*http.Response, error) {
	if data != nil && data.Get("authenticity_token") == "" {
		data.Add("authenticity_token", s.CSRF)
	}

	req, err := http.NewRequest(http.MethodPost, caviarURL(url), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	return s.Client.Do(req)
}

// Get behaves *http.Client.Get while automatically setting the URL prefix and
// CSRF token.
func (s *Session) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, caviarURL(url), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CSRF-Token", s.CSRF)
	req.Header.Add("Accept", "application/json")

	return s.Client.Do(req)
}
