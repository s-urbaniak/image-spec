package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/opencontainers/image-spec/oci/statik"
)

func main() {
	cmd := &cobra.Command{
		Use:   "oci-image-tool",
		Short: "A tool for working with OCI images",
	}

	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)

	cmd.AddCommand(newValidateCmd(stdout, stderr))
	cmd.Execute()
}
