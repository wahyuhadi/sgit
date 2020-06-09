package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v32/github"
)

func GetListCommits(session *http.Client, owner, repo string) {
	until := time.Now()
	since := until.AddDate(-1, 0, 0)
	client := github.NewClient(session)
	page := 1
	for true {
		opt := &github.CommitsListOptions{
			Since:       since,
			Until:       until,
			ListOptions: github.ListOptions{Page: page, PerPage: 100},
		}
		commits, resp, err := client.Repositories.ListCommits(context.Background(), owner, repo, opt)
		if err != nil {
			fmt.Println(errors.New("[+] Error when get list commits"))
		}

		for _, isDataCommits := range commits {
			fmt.Println("\t[+] Scan Sha commits :", *isDataCommits.SHA)
			GetCommit(session, owner, repo, *isDataCommits.SHA)
		}

		isNextPage := resp.NextPage
		if isNextPage == 0 {
			break
		}
		page = isNextPage
	}
}

func GetCommit(session *http.Client, owner, repo, sha string) {
	client := github.NewClient(session)
	commits, _, err := client.Repositories.GetCommit(context.Background(), owner, repo, sha)
	if err != nil {
		fmt.Println(errors.New("[+] Error when get commits"))
	}

	fmt.Println("\t\t[+] Status ", commits.Stats.Total)
	for _, isData := range commits.Files {
		if *isData.Filename != ".DS_Store" {
			fmt.Println("\t\t[+] File Name", *isData.Filename)
			fmt.Println("\t\t[+] Patch ", *isData.Patch)
			fmt.Println("\t\t[+] File ", repo)
		}
	}
}
