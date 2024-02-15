package src

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/provider.tf.gotpl
//go:embed templates/terraform.tfvars.gotpl
//go:embed templates/op-compute-restrictSharedVpcSubnetworks.tf.gotpl
var providerTmpl embed.FS

type tfvars struct {
	Project string
}

type subnets struct {
	ProjectName string
	ProjectID   string
	Region      string
	SubnetName  string
}

func createAndExecuteTemplate(templateName string, outputName string, dir string, data interface{}) {
	fmt.Println("Creating file:", outputName)
	tmpl, err := template.ParseFS(providerTmpl, templateName)
	if err != nil {
		log.Println("Error parsing template:", templateName, err)
		return
	}
	out, err := os.Create(filepath.Join(dir, outputName))
	if err != nil {
		log.Println("Error creating file:", outputName, err)
		return
	}
	defer out.Close()
	err = tmpl.Execute(out, data)
	if err != nil {
		log.Println("Error executing template:", templateName, err)
		return
	}
	fmt.Println("File created:", outputName)
}

func CreateTF(dir string, project string) {
	createAndExecuteTemplate("templates/provider.tf.gotpl", "provider.tf", dir, nil)
	createAndExecuteTemplate("templates/terraform.tfvars.gotpl", "terraform.tfvars", dir, tfvars{Project: project})
	createAndExecuteTemplate("templates/op-compute-restrictSharedVpcSubnetworks.tf.gotpl", "op-compute-restrictSharedVpcSubnetworks.tf", dir, subnets{ProjectName: project, ProjectID: "lab-netops", Region: "north1", SubnetName: "subnet-1"})
}
