package cmds

import (
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateVariables(cloudprovider string) (string, error) {
	templatepath := "templates/" + cloudprovider + "/createvariables.tpl"
	tpl, err := pongo2.FromFile(templatepath)

	print(tpl, "the file")
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	// defer awaits the function to finish before executing

	var terraformStr strings.Builder
	err = tpl.ExecuteWriter(pongo2.Context{}, &terraformStr)
	if err != nil {
		panic(err)
	}
	return terraformStr.String(), nil
}

// func CreateProvider(provider pkgs.Prov) error {
// 	tpl, err := pongo2.FromFile("templates/aws/createprovider.tpl")
// 	if err != nil {
// 		return err
// 	}

// 	file, err := os.OpenFile("terraform/provider.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	fmt.Println("Writing to file: terraform/provider.tf")
// 	err = tpl.ExecuteWriter(pongo2.Context{"provider": provider}, file)
// 	return err
// }

func CreateProvider(provider pkgs.Prov) (string, error) {
	templatepath := "templates/" + provider.Provider + "/createprovider.tpl"

	// tpl, err := pongo2.FromFile("templates/aws/createprovider.tpl")
	tpl, err := pongo2.FromFile(templatepath)
	if err != nil {
		return "", err
	}

	var terraformStr strings.Builder

	err = tpl.ExecuteWriter(pongo2.Context{"provider": provider}, &terraformStr)
	if err != nil {
		return "", err
	}
	formatted := hclwrite.Format([]byte(terraformStr.String()))

	return string(formatted), nil
}
