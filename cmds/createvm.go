package cmds

import (
	"generatorv/pkgs"

	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateVirtualMachine(vms []pkgs.VirtualMachine, provider string) (string, error) {
	templatepath := "templates/" + provider + "/compute/vm.tpl"

	tpl, err := pongo2.FromFile(templatepath)

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
