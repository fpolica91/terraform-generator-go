package cmds

import (
	"generatorv/pkgs"

	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateVirtualMachine(vms []pkgs.VirtualMachine, provider string, compute_type string) (string, error) {

	template_path := "templates/" + provider + "/compute/" + compute_type + "/create.tpl"
	tpl, err := pongo2.FromFile(template_path)

	if err != nil {
		return "", err
	}

	var resourceString strings.Builder

	err = tpl.ExecuteWriter(pongo2.Context{"vms": vms}, &resourceString)
	if err != nil {
		return "", err
	}
	formatted := hclwrite.Format([]byte(resourceString.String()))

	return string(formatted), nil
}
