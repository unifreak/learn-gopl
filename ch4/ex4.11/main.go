// @todo

// Build a tool that lets users create, read, update, and delete GitHub issues
// from the command line, invoking their preferred text editor when substantial
// text input is required.

// Github Rest Api Doc:
//
// Authentication: curl -i -u your_username:$token
//
// List all your issues:
//
// 		curl -i -H "Authorization: token ghp_16C7e42F292c6912E7710c838347Ae178B4a" \
//      https://api.github.com/issues
//
// Create:
//
//  	curl -i -H 'Authorization: token ghp_16C7e42F292c6912E7710c838347Ae178B4a' \
// 		   -d '{ \
// 		        "title": "New logo", \
// 		        "body": "We should have one", \
// 		        "labels": ["design"] \
// 		      }' \
// 		   https://api.github.com/repos/OWNER/REPO/issues
//
// Get:
//
// 		curl \
// 		  -H "Accept: application/vnd.github.v3+json" \
// 		  https://api.github.com/repos/OWNER/REPO/issues/ISSUE_NUMBER
//
// Update:
//
// 		curl \
// 			-X PATCH \
// 			-H "Accept: application/vnd.github.v3+json" \
// 			https://api.github.com/repos/OWNER/REPO/issues/ISSUE_NUMBER \
// 			-d '{"title":"Found a bug","body":"I'm having a problem ...}'
//

package github

import (
	"fmt"
	. "learn/gopl/ch4/github"
	"net/http"
	"os"
)

const (
	ListURL = "https://api.github.com/issues"
)

// @?
// - how to handle token?

type IssueList []*Issue

func ListIssues() (*IssueList, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", ListURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authentication", "token "+os.Getenv("GHTOKEN"))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", resp.Body)
}

func main() {
	ListIssues()
}