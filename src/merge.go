package src

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func CreateMR(token string, project gitlab.Project, client *gitlab.Client) {

	opt := &gitlab.CreateMergeRequestOptions{
		Title:        gitlab.Ptr("Generating Terraform subnets"),
		Description:  gitlab.Ptr("Merge Request"),
		SourceBranch: gitlab.Ptr("terraform-subnets"),
		TargetBranch: gitlab.Ptr("develop"),
	}

	repoID := project.ID

	fmt.Println("Creating merge request")

	mr, _, err := client.MergeRequests.CreateMergeRequest(repoID, opt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created new merge request: %s \n", mr.Title)
}
