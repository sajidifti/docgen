//go:generate go run generate-asset.go

package main

import (
	"github.com/sajidifti/docgen/cmd"
)

func main() {
	cmd.Execute()
}
