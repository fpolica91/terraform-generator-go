package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"os"

	"github.com/flosch/pongo2/v6"
)

func CreateVirtualPrivateCloud(vpcs []pkgs.VPC) error {
	tpl, err := pongo2.FromFile("templates/aws/vpc/createvpc.tpl")
	if err != nil {
		return err
	}

	file, err := os.OpenFile("terraform/vpc.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Writing to file: terraform/vpc.tf")

	err = tpl.ExecuteWriter(pongo2.Context{"vpcs": vpcs}, file)
	return err // or handle it as you prefer
}
