package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"os"

	"github.com/flosch/pongo2/v6"
)

func CreateVariables() error {
	tpl, err := pongo2.FromFile("templates/aws/createvariables.tpl")

	print(tpl, "the file")
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("terraform/variables.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	// defer awaits the function to finish before executing
	defer file.Close()
	fmt.Println("Writing to file: terraform/variables.tf")
	err = tpl.ExecuteWriter(pongo2.Context{}, file)

	return err
}

func CreateProvider(provider pkgs.Prov) error {
	tpl, err := pongo2.FromFile("templates/aws/createprovider.tpl")
	if err != nil {
		return err
	}

	file, err := os.OpenFile("terraform/provider.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Writing to file: terraform/provider.tf")
	err = tpl.ExecuteWriter(pongo2.Context{"provider": provider}, file)
	return err
}
