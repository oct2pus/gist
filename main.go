package main

import (
	"context"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-github/v32/github"
)

func main() {
	userPtr := flag.String("u", "", "specifies user to list.")
	/*	config := &oauth2.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{"gist"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  github.URL,
				TokenURL: "",
			},
		}
		oauth2.NewClient()*/
	flag.Parse()
	c := github.NewClient(nil)
	gists, _, _ := c.Gists.List(context.Background(), *userPtr, &github.GistListOptions{})
	for _, gist := range gists {
		fmt.Printf("%v/%v - %v file(s)\n%v\n%v\n\n", gist.GetOwner().GetLogin(), getFilenames(gist.Files)[0], len(gist.Files), gist.GetUpdatedAt(), gist.GetDescription())
	}
}

func getFilenames(files map[github.GistFilename]github.GistFile) []string {
	if len(files) == 1 {
		for name := range files {
			return []string{string(name)} // this is stupid
		}
	} else {
		names := make([]string, len(files), len(files))
		i := 0
		for name := range files {
			names[i] = string(name)
			i++
		}
		sort.Strings(names)
		return names
	}
	return []string{""}
}

func listFiles(files map[github.GistFilename]github.GistFile) string {
	var build strings.Builder
	for _, ele := range files {
		build.WriteString(ele.GetFilename())
		build.WriteString(" - ")
		build.WriteString(ele.GetLanguage())
		build.WriteRune('\n')
	}
	return build.String()
}
