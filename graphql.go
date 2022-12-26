package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const endpoint = "https://api.github.com/graphql"

var limit = wf.Config.Get("limit")

type SearchResult struct {
	Title    string
	Subtitle string
	Arg      string
}

// generated using https://transform.tools/json-to-go
type issueSearchResult struct {
	Data struct {
		Search struct {
			Nodes []struct {
				Title      string `json:"title"`
				URL        string `json:"url"`
				CreatedAt  string `json:"createdAt"`
				Repository struct {
					Name string `json:"name"`
				} `json:"repository"`
				Author struct {
					Login string `json:"login"`
				} `json:"author"`
			} `json:"nodes"`
		} `json:"search"`
	} `json:"data"`
}

// generated using https://transform.tools/json-to-go
type repoSearchResult struct {
	Data struct {
		Search struct {
			Nodes []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nodes"`
		} `json:"search"`
	} `json:"data"`
}

// search searches issues or repositories using GitHub's GraphQL endpoint
// for a given query.
func Search(q string) []*SearchResult {
	typeAndQuery := strings.Split(q, "§")

	switch typeAndQuery[0] {
	case "ISSUE":
		return searchIssues(typeAndQuery[1])
	case "REPOSITORY":
		return searchRepos(typeAndQuery[1])
	default:
		return []*SearchResult{}
	}
}

// searchIssues builds a query for searching issues.
// Results are deserialized and returned as a SearchResult.
func searchIssues(q string) []*SearchResult {
	var gqlFormat = `{"query":"query{search(query:\"%s\",type:ISSUE,last:%s){nodes{...on Issue{title url createdAt author{login}repository{name}}...on PullRequest{title url createdAt author{login}repository{name}}}}}"`
	q = strings.Replace(q, "ISSUE", "", 1)
	bodyBytes := query(fmt.Sprintf(gqlFormat, q, limit))

	// deserialize response to struct
	var issues issueSearchResult
	json.Unmarshal(bodyBytes, &issues)

	var results []*SearchResult
	for _, issue := range issues.Data.Search.Nodes {
		results = append(
			results,
			&SearchResult{
				issue.Title,
				issue.Author.Login + " | " + issue.CreatedAt + " | " + issue.Repository.Name,
				issue.URL,
			},
		)
	}

	return results
}

// searchRepos builds a query for searching repositories.
// Results are deserialized and returned as a SearchResult.
func searchRepos(q string) []*SearchResult {
	var gqlFormat = `{"query":"query{search(query:\"%s\",type:REPOSITORY,last:%s){nodes{...on Repository{name url}}}}"`
	q = strings.Replace(q, "REPOSITORY", "", 1)
	bodyBytes := query(fmt.Sprintf(gqlFormat, q, limit))

	// deserialize response to struct
	var repos repoSearchResult
	json.Unmarshal(bodyBytes, &repos)

	var results []*SearchResult
	for _, repo := range repos.Data.Search.Nodes {
		results = append(
			results,
			&SearchResult{
				repo.Name,
				repo.URL,
				repo.URL,
			},
		)
	}

	return results
}

// query makes a post request to GitHub's GraphQL endpoint
// for a given query, using a cached or new oauth token.
func query(q string) []byte {
	// get cached token
	tok, err := CachedToken()
	if err != nil {
		log.Printf("Error retrieving cached token; it might not exist: %v", err)

		// get new token
		tok, err = newToken()
		if err != nil {
			log.Fatalf("Error aquiring token: %v", err)
		}

		// store token
		err = CacheToken(tok)
		if err != nil {
			log.Fatalf("Error storing token: %v", err)
		}
	}

	// build request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(q)))
	if err != nil {
		log.Fatalf("Error building request: %v", err)
	}
	req.Header.Set("Authorization", "bearer "+tok.AccessToken)

	// make request
	var client = http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Response unsuccessful: %v", resp.Status)
	}
	defer resp.Body.Close()

	// read response as bytes
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return bodyBytes
}
