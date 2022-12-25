package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const endpoint = "https://api.github.com/graphql"

type SearchResult struct {
	Data struct {
		Search struct {
			Nodes []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nodes"`
		} `json:"search"`
	} `json:"data"`
}

func search(q string) SearchResult {
	var searchRepos = `{"query":"query{search(query:\"%s\",type:REPOSITORY,last:100){nodes{...on Repository{name url}}}}"`
	return query(fmt.Sprintf(searchRepos, q))
}

// query makes a post request to GitHub's GraphQL endpoint
// for a given query, using a cached or new oauth token.
func query(q string) SearchResult {
	// get cached token
	tok, err := cachedToken()
	if err != nil {
		log.Printf("Error retrieving cached token; it might not exist: %v", err)

		// get new token
		tok, err = newToken()
		if err != nil {
			log.Fatalf("Error aquiring token: %v", err)
		}

		// store token
		err = cacheToken(tok)
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

	// deserialize response to struct
	var results SearchResult
	json.Unmarshal(bodyBytes, &results)

	return results
}
