package src

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func CreateBranch(token string, project gitlab.Project) {
	git, _ := gitlab.NewClient(token)

	opt := &gitlab.CreateBranchOptions{
		Branch: gitlab.Ptr("develop"),
		Ref:    gitlab.Ptr("master"),
	}

	repoID := project.ID

	branch, _, err := git.Branches.CreateBranch(repoID, opt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created new branch: %s \n", branch.Name)
}
