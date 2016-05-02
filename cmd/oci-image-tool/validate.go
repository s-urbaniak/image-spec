package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/opencontainers/image-spec/oci"
	"github.com/opencontainers/image-spec/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	typeImageLayout  = "imageLayout"
	typeImage        = "image"
	typeManifest     = "manifest"
	typeManifestList = "manifestList"
	typeConfig       = "config"
)

var validateTypes = []string{
	typeImageLayout,
	typeImage,
	typeManifest,
	typeManifestList,
	typeConfig,
}

type validateCmd struct {
	stdout *log.Logger
	stderr *log.Logger
	typ    string // the type to validate, can be empty string
}

func newValidateCmd(stdout, stderr *log.Logger) *cobra.Command {
	v := &validateCmd{
		stdout: stdout,
		stderr: stderr,
	}

	cmd := &cobra.Command{
		Use:   "validate FILE...",
		Short: "Validate one or more image files",
		Run:   v.Run,
	}

	cmd.Flags().StringVar(
		&v.typ, "type", "",
		fmt.Sprintf(
			`Type of the file to validate. If unset, oci-image-tool will try to auto-detect the type. One of "%s"`,
			strings.Join(validateTypes, ","),
		),
	)

	return cmd
}

func (v *validateCmd) Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		v.stderr.Printf("no files specified")
		cmd.Usage()
		os.Exit(1)
	}

	var exitcode int
	for _, arg := range args {
		err := v.validatePath(arg)

		if err == nil {
			continue
		}

		var verr oci.ValidationError
		var ok bool
		if verr, ok = errors.Cause(err).(oci.ValidationError); !ok {
			v.stderr.Println(err)
			exitcode = 1
			continue
		}

		for _, err := range verr.Errs {
			v.stderr.Printf("path %q: validation failed: %v", arg, err)
		}

		exitcode = 1
	}

	if exitcode == 0 {
		v.stdout.Println("OK")
	}

	os.Exit(exitcode)
}

func (v *validateCmd) validatePath(name string) error {
	var err error
	typ := v.typ

	if typ == "" {
		if typ, err = autodetect(name); err != nil {
			return errors.Wrap(err, "unable to determine type")
		}
	}

	switch typ {
	case typeManifest:
		f, err := os.Open(name)
		if err != nil {
			errors.Wrap(err, "unable to open file")
		}
		defer f.Close()

		if err := oci.ValidateManifest(f); err != nil {
			return err
		}

		return nil
	case typeManifestList:
		f, err := os.Open(name)
		if err != nil {
			errors.Wrap(err, "unable to open file")
		}
		defer f.Close()

		if err := oci.ValidateManifestList(f); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("path %q: type %q unimplemented", name, typ)
}

// autodetect returns one of the type.. constants of the content of the given path
// or an error if the type could not be resolved.
func autodetect(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", errors.Wrapf(err, "unable to access path") // err from os.Stat includes path name
	}

	if fi.IsDir() {
		return typeImageLayout, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "unable to open file") // os.Open includes the filename
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(io.LimitReader(f, 512)) // read up to 512 bytes for content detection
	if err != nil {
		return "", errors.Wrapf(err, "unable to read file %q", path)
	}

	mimeType := http.DetectContentType(buf)

	switch mimeType {
	case "application/x-gzip":
		return typeImage, nil

	case "application/octet-stream":
		return typeImage, nil

	case "text/plain; charset=utf-8":
		// might be a JSON file, will be handled below

	default:
		return "", fmt.Errorf("unknown file type: file %q", path)
	}

	if _, err := f.Seek(0, os.SEEK_SET); err != nil {
		return "", errors.Wrapf(err, "unable to read file %q", path)
	}

	header := struct {
		SchemaVersion int         `json:"schemaVersion"`
		MediaType     string      `json:"mediaType"`
		Config        interface{} `json:"config"`
	}{}

	if err := json.NewDecoder(f).Decode(&header); err != nil {
		return "", errors.Wrapf(err, "unable to parse JSON file: %q", path)
	}

	switch {
	case header.MediaType == schema.MediaTypeManifest:
		return typeManifest, nil

	case header.MediaType == schema.MediaTypeManifestList:
		return typeManifestList, nil

	case header.MediaType == "" && header.SchemaVersion == 0 && header.Config != nil:
		// config files don't have mediaType/schemaVersion header
		return typeConfig, nil
	}

	return "", fmt.Errorf("unknown media type: file %q", path)
}
