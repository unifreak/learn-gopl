// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// See https://developer.github.com/v3/
const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	// the matching process that associates JSON names with Go struct names
	// during unmarshaling is case-insensitive, sot it's only neccesary to
	// use a field tag when there's an underscore in the JSON name but not
	// in the Go name.
	TotalCount 	int `json:"total_count"`
	Items 		[]*Issue
}

type Issue struct {
	Number 		int
	HTMLURL 	string `json:"html_url"`
	Title 		string
	State 		string
	User 		*User
	CreatedAt 	time.Time `json:"created_at"`
	Body 		string 	// in MarkDown format
}

type User struct {
	Login 		string
	HTMLURL 	string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	// Unlike in movie/, we use the `streaming decoder` json.Decoder here,
	// which allows several JSON entities to be decoded in sequence from the
	// same stream (we don't need that feature here).
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}