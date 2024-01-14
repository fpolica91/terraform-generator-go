package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateObjectStorage(storage_objects []pkgs.ObjectStorage, provider string, storagetype string) (string, error) {
	template_path := "templates/" + provider + "/storage/" + storagetype + "/create.tpl"
	fmt.Println("Using template: " + template_path)
	fmt.Println(storage_objects, "storage_objects")

	tpl, err := pongo2.FromFile(template_path)
	if err != nil {
		panic(err)

	}

	var terraformStr strings.Builder
	err = tpl.ExecuteWriter(pongo2.Context{"storage_objects": storage_objects}, &terraformStr)
	if err != nil {
		panic(err)

	}

	formattedstring := hclwrite.Format([]byte(terraformStr.String()))

	return string(formattedstring), nil
}
