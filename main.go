package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	client := getClient(ctx)
	repos, _, err := client.Repositories.List(ctx, "borankux", &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	if err != nil {
		panic(err)
	}
	for idx, repo := range repos {
		fmt.Println(fmt.Sprintf("#%s, %s, %s, \n %s \n\n",
			color.WhiteString("%d", idx),
			color.CyanString("%s", repo.GetLanguage()),
			color.GreenString("%s", repo.GetFullName()),
			color.BlueString("%s", repo.GetDescription()),
		))
	}
}

func getClient(ctx context.Context) *github.Client{
	token := getConfig().GetString("key")
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken:  token,
	})

	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func getConfig() *viper.Viper{
	v := viper.GetViper()
	v.AddConfigPath("./")
	v.SetConfigName("credentials")
	v.SetConfigType("yml")
	err := v.ReadInConfig()

	if err != nil {
		panic(v)
	}
	return v
}