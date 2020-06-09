package main

import (
	"context"
	"flag"
	"os"
	"sgit/services"

	"golang.org/x/oauth2"
)

var (
	org  = flag.String("org", "github", "Your Org ")
	slug = flag.String("slug", "slug", "Your Slug ")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("sgit")},
	)

	tc := oauth2.NewClient(ctx, ts)
	_, _ = services.GetList(tc, *org, *slug)
}
