package src

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func CreateBranch(token string, project gitlab.Project, client *gitlab.Client) {

	opt := &gitlab.CreateBranchOptions{
		Branch: gitlab.Ptr("develop"),
		Ref:    gitlab.Ptr("master"),
	}

	repoID := project.ID

	branch, _, err := client.Branches.CreateBranch(repoID, opt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created new branch: %s \n", branch.Name)
}
