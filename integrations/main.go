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
	client := github.NewClient(nil)

	// list all organizations for user "harshitha-akkaraju"
	orgs, _, err := client.Organizations.List(context.Background(), "harshitha-akkaraju", nil)
	// orgs, _, err := client.Organizations.List(context.Background(), "LiorLev", nil)

	// orgLogins := make([]string, len(orgs))
	// for _, o := range orgs {
	// 	orgLogins = append(orgLogins, o.GetLogin())
	// }
	// fmt.Println(orgLogins)

	fmt.Println(len(orgs))
	// fmt.Println(orgs)

	// Get repos for authenticated user
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	clientInstance2 := github.NewClient(tc)

	teams, response, err := clientInstance2.Teams.ListUserTeams(ctx, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(response)
	}

	for _, t := range teams {
		reposURL := t.GetRepositoriesURL()
		request, _ := clientInstance2.NewRequest(http.MethodGet, reposURL, nil)
		repos := make([]*repository, 5)
		response, err := clientInstance2.Do(ctx, request, &repos)
		if err != nil {
			fmt.Println(err)
			fmt.Println(response)
		}
		// TODO: unmarshall into Repo object
		for _, repo := range repos {
			fmt.Println(repo.FullName)
		}
	}

	// list all repositories for the authenticated user
	repos, _, err := clientInstance2.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(repos))
	// fmt.Println(repos)

	// Get teams for a user
	teamList, _, err := clientInstance2.Teams.ListTeams(ctx, "fabfourfems", nil)
	fmt.Println(teamList)
}
