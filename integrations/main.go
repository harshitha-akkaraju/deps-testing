package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github" // with go modules disabled
	"golang.org/x/oauth2"
)

type repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

func main() {
	// Get repos for authenticated user
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	teams, response, err := client.Teams.ListUserTeams(ctx, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(response)
	}

	authUserRepos := make([]*repository, 0)
	for _, t := range teams {
		reposURL := t.GetRepositoriesURL()
		request, _ := client.NewRequest(http.MethodGet, reposURL, nil)
		repos := make([]*repository, 0)
		response, err := client.Do(ctx, request, &repos)
		if err != nil {
			fmt.Println(err)
			fmt.Println(response)
		}
		authUserRepos = append(authUserRepos, repos...)
	}

	for _, repo := range authUserRepos {
		fmt.Println(repo.FullName)
	}
}
