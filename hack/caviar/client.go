package caviar

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// caviarURL joins a URL with the Caviar prefix.
func caviarURL(url string) string {
	return URL + "/" + strings.TrimPrefix(url, "/")
}

// PostForm makes a URL-encoded form POST with the correct base URL and CSRF
// token.
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

// PostJSON makes a JSON POST with the correct base URL and CSRF token.
func (s *Session) PostJSON(url string, data interface{}) (*http.Response, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, caviarURL(url), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CSRF-Token", s.CSRF)
	req.Header.Add("Accept", "application/json")

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Println(string(dump))

	return s.Client.Do(req)
}

// GetJSON makes a GET request with the correct base URL, content type, and CSRF
// token.
func (s *Session) GetJSON(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, caviarURL(url), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CSRF-Token", s.CSRF)
	req.Header.Add("Accept", "application/json")

	return s.Client.Do(req)
}
