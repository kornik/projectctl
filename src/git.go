package src

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/xanzy/go-gitlab"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func CloneCommitPush(repoURL gitlab.Project, token string, client *gitlab.Client) {
	// Clone the repository

	dir, err := os.MkdirTemp("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	// clean up

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:        repoURL.HTTPURLToRepo,
		RemoteName: "develop",
		Auth: &http.BasicAuth{
			Username: "abc", // this can be anything except an empty string
			Password: token,
		},
	})
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	if err != nil {
		log.Fatalf("Could not clone the repository: %s", err)
	}
	//re, err := git.PlainOpen(dir)

	// Open the worktree
	w, err := r.Worktree()
	if err != nil {
		log.Fatalf("Could not access the worktree: %s", err)
	}

	b := fmt.Sprint("refs/heads/terraform-subnets")
	branch := plumbing.ReferenceName(b)
	err = w.Checkout(&git.CheckoutOptions{
		Create: true,
		Branch: branch,
	})
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: branch.Short(),
		URLs: []string{repoURL.HTTPURLToRepo},
	})
	if err != nil {
		log.Fatalf("Could not create the remote: %s", err)
	}

	fmt.Println(w.Status())
	err = w.Checkout(&git.CheckoutOptions{
		Branch: branch,
	})
	CreateTF(dir, repoURL.Name)

	// Add the new file to the project
	_, err = w.Add(".")
	if err != nil {
		log.Fatalf("Could not add the file to the project: %s", err)
	}

	// Commit the changes
	_, err = w.Commit("Add file via go-git", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "go-git User",
			Email: "gogit@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("Could not commit the changes: %s", err)
	}
	fmt.Println("Commit done")

	// Push the changes
	err = r.Push(&git.PushOptions{
		RemoteName: branch.Short(),

		Auth: &http.BasicAuth{
			Username: "abc", // This can be anything except an empty string.
			Password: token,
		},
	})
	if err != nil {
		log.Printf("Something went wrong at push: %s\n", err)
		log.Println("Deleting project...")
		DeleteProject(token, repoURL, client)
		log.Println("Project deleted")
	}
	CreateMR(token, repoURL)

	defer os.RemoveAll(dir)
}
