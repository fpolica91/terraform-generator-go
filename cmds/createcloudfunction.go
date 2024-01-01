package cmds

import (
	"fmt"
	"generatorv/pkgs"

	"strings"

	"github.com/flosch/pongo2/v6"
)

func CreateCloudFunction(cloudfuction []pkgs.CloudFunction, provider string) (string, error) {
	templatepath := "templates/" + provider + "/compute/function.tpl"
	fmt.Println(cloudfuction, "cloudfuction")
	fmt.Println(provider, "provider")

	tpl, err := pongo2.FromFile(templatepath)

	if err != nil {
		return "", err
	}

	var resourceString strings.Builder

	err = tpl.ExecuteWriter(pongo2.Context{"lambdas": cloudfuction}, &resourceString)
	if err != nil {
		return "", err
	}
	fmt.Println(resourceString.String(), "resourceString")
	return resourceString.String(), nil
}
