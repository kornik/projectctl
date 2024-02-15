package src

import "github.com/xanzy/go-gitlab"

func CreateClient(token string, gitlabUrl string) *gitlab.Client {

	git, _ := gitlab.NewClient(token, gitlab.WithBaseURL(gitlabUrl))
	return git
}
