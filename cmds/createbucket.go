package cmds

import (
	"generatorv/pkgs"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// func CreateBuckets(buckets []pkgs.Bucket) error {
// 	tpl, err := pongo2.FromFile("templates/aws/s3/createbucket.tpl")
// 	if err != nil {
// 		return err
// 	}

// 	file, err := os.OpenFile("terraform/s3.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	fmt.Println("Writing to file: terraform/s3.tf")

// 	err = tpl.ExecuteWriter(pongo2.Context{"buckets": buckets}, file)
// 	return err // or handle it as you prefer
// }

func CreateBuckets(buckets []pkgs.Bucket) (string, error) {
	tpl, err := pongo2.FromFile("templates/aws/s3/createbucket.tpl")
	if err != nil {
		return "", err
	}

	var terraformStr strings.Builder
	err = tpl.ExecuteWriter(pongo2.Context{"buckets": buckets}, &terraformStr)
	if err != nil {
		return "", err
	}

	formattedstring := hclwrite.Format([]byte(terraformStr.String()))

	return string(formattedstring), nil
}
