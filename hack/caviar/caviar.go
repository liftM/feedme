package caviar

import (
	"errors"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

const (
	// URL is the prefix for Caviar requests.
	URL = "https://www.trycaviar.com"
)

// Session contains a Caviar web application session state.
type Session struct {
	CSRF string

	*http.Client
}

// New creates and begins a new session.
func New() (*Session, error) {
	// Create a new client that maintains a cookie jar.
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	client := &http.Client{
		Jar: jar,
	}

	// Load sign-in page.
	res, err := client.Get(URL + "/users/sign_in")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	// Get CSRF token.
	sel := doc.Find("meta[name='csrf-token']")
	if sel.Length() != 1 {
		return nil, errors.New("no CSRF token found")
	}

	node := sel.Get(0)
	var token string
	for _, attr := range node.Attr {
		if attr.Key == "name" && attr.Val != "csrf-token" {
			return nil, errors.New("invalid CSRF token")
		}
		if attr.Key == "content" {
			token = attr.Val
		}
	}

	return &Session{
		CSRF:   token,
		Client: client,
	}, nil
}
