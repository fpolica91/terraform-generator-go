package cmds

import (
	"fmt"
	"generatorv/pkgs"
	"os"

	"github.com/flosch/pongo2/v6"
)

func CreateBuckets(buckets []pkgs.Bucket) error {
	tpl, err := pongo2.FromFile("templates/aws/s3/createbucket.tpl")
	if err != nil {
		return err
	}

	file, err := os.OpenFile("terraform/s3.tf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Writing to file: terraform/s3.tf")

	err = tpl.ExecuteWriter(pongo2.Context{"buckets": buckets}, file)
	return err // or handle it as you prefer
}
