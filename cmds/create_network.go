package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateNetwork(vpcs []pkgs.NetworkUnit, provider string, cloud_type string) (string, error) {
	template_path := "templates/" + provider + "/network/" + cloud_type + "/create.tpl"
	fmt.Println("Using template: " + template_path)
	tpl, err := pongo2.FromFile(template_path)
	if err != nil {
		return "", err
	}

	fmt.Println("Writing to file: terraform/vpc.tf")
	var resourcesString strings.Builder

	err = tpl.ExecuteWriter(pongo2.Context{"vpcs": vpcs}, &resourcesString)
	if err != nil {
		return "", err
	}

	formattedString := hclwrite.Format([]byte(resourcesString.String()))

	return string(formattedString), nil

}
