package src

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func CreateProject(token string, name string, description string, client *gitlab.Client) *gitlab.Project {

	opt := &gitlab.CreateProjectOptions{
		Name:                 gitlab.Ptr(name),
		Description:          gitlab.Ptr(description),
		InitializeWithReadme: gitlab.Ptr(true),
		DefaultBranch:        gitlab.Ptr("master"),
	}

	project, _, err := client.Projects.CreateProject(opt)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Project created: ", project.Name)
	return project
}

func DeleteProject(token string, project gitlab.Project, client *gitlab.Client) {

	_, err := client.Projects.DeleteProject(project.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Project deleted: ", project.Name)
}
