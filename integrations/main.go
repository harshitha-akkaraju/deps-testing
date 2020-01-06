package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github" // with go modules disabled
	"golang.org/x/oauth2"
)

func main() {
	client := github.NewClient(nil)

	// list all organizations for user "harshitha-akkaraju"
	orgs, _, err := client.Organizations.List(context.Background(), "harshitha-akkaraju", nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(orgs))

	// Get repos for authenticated user
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubAccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	clientInstance2 := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := clientInstance2.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(repos))

	// Get teams for a user
	teamList, _, err := clientInstance2.Teams.ListTeams(ctx, "fabfourfems", nil)
	fmt.Println(teamList)
}
