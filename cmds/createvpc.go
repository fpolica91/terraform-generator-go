package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateVirtualPrivateCloud(vpcs []pkgs.VPC) (string, error) {
	tpl, err := pongo2.FromFile("templates/aws/vpc/createvpc.tpl")
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
