package cmds

import (
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func CreateBuckets(buckets []pkgs.Bucket, provider string, storagetype string) (string, error) {
	template_path := "templates/" + provider + "/storage/" + storagetype + "/create.tpl"

	tpl, err := pongo2.FromFile(template_path)
	if err != nil {
		panic(err)

	}

	var terraformStr strings.Builder
	err = tpl.ExecuteWriter(pongo2.Context{"buckets": buckets}, &terraformStr)
	if err != nil {
		panic(err)

	}

	formattedstring := hclwrite.Format([]byte(terraformStr.String()))

	return string(formattedstring), nil
}
