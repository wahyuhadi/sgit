package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v32/github"
	// with go modules disabled
)

func GetList(session *http.Client, org, slug string) ([]string, error) {
	client := github.NewClient(session)
	var RepositoryName []string
	page := 1

	for true {
		opt := &github.ListOptions{Page: page, PerPage: 100}

		// Get Repos by Org - and Slug
		teams, resp, err := client.Teams.ListTeamReposBySlug(context.Background(), org, slug, opt)
		if err != nil {
			return nil, errors.New("[!] error when get list repo")
		}

		for _, repo := range teams {
			fmt.Println("[+] Repos Owner : ", *repo.Owner.Login)
			fmt.Println("[+] Repos Name : ", *repo.Name)
			fmt.Println("[+] Repo URL : ", *repo.HTMLURL)
			GetListCommits(session, *repo.Owner.Login, *repo.Name)
		}
		// get NextPage
		isNextPage := resp.NextPage
		if isNextPage == 0 {
			break
		}

		// Adder page
		page = isNextPage
	}

	return RepositoryName, nil
}
