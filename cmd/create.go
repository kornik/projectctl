/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"projectctl/src"
)

var name string
var token string
var gitlabUrl string
var client = src.CreateClient(token, gitlabUrl)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//src.CreateTF(".", token)
		project := src.CreateProject(token, name, "test", client)
		src.CreateBranch(token, *project)
		fmt.Println(project.HTTPURLToRepo)
		src.CloneCommitPush(*project, token)
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the project")
	createCmd.MarkPersistentFlagRequired("name")
	createCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Gitlab token")
	createCmd.MarkPersistentFlagRequired("token")
	createCmd.PersistentFlags().StringVarP(&gitlabUrl, "url", "u", "", "Gitlab url")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
