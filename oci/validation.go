package oci

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/opencontainers/image-spec/schema"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/s-urbaniak/gojsonschema"

	_ "github.com/opencontainers/image-spec/oci/statik"
)

var (
	specs = map[string]string{
		schema.MediaTypeManifest:     "image-manifest-schema.json",
		schema.MediaTypeManifestList: "manifest-list-schema.json",
	}
)

type ValidationError struct {
	Errs []error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%v", e.Errs)
}

func ValidateManifest(src io.Reader) error {
	return validate(src, specs[schema.MediaTypeManifest])
}

func ValidateManifestList(src io.Reader) error {
	return validate(src, specs[schema.MediaTypeManifestList])
}

func validate(src io.Reader, schema string) error {
	statikFS, err := fs.New()
	if err != nil {
		return errors.Wrap(err, "unable to initialize statik filesystem")
	}

	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return errors.Wrap(err, "unable to read manifest")
	}

	sl := gojsonschema.NewReferenceLoaderFileSystem("file:///"+schema, statikFS)
	ml := gojsonschema.NewStringLoader(string(buf))

	result, err := gojsonschema.Validate(sl, ml)
	if err != nil {
		return errors.Wrap(err, "unable to validate manifest")
	}

	if result.Valid() {
		return nil
	}

	errs := make([]error, 0, len(result.Errors()))
	for _, desc := range result.Errors() {
		errs = append(errs, fmt.Errorf("%s", desc))
	}

	return ValidationError{
		Errs: errs,
	}
}
