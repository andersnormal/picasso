package providers

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/andersnormal/picasso/pkg"
	"github.com/andersnormal/picasso/config"
	"github.com/andersnormal/picasso/pkg/spec"

	"gopkg.in/yaml.v2"
)

type archiveProvider struct {
	cfg *config.Config
}

// NewArchive ...
func NewArchive(cfg *config.Config) pkg.Provider {
	p := new(archiveProvider)
	p.cfg = cfg

	return p
}

// CloneWithContext ...
func (a *archiveProvider) CloneWithContext(ctx context.Context) error {
	resp, err := http.Get(a.cfg.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// reading the zip from the body
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	// this gives the files in the zip in a slice
	var assets []*zip.File
	var y *zip.File

	// stream all the files ... maybe copy to the filesystem
	for _, f := range zipReader.File {
		if !strings.Contains(f.Name, ".sc.yaml") {
			assets = append(assets, f)

			continue
		}

		y = f
	}

	if y == nil {
		return errors.New(`no sc spec found`)
	}

	// this should be later filtered to be the root of the files
	base := path.Base(y.Name)

	var s *spec.Spec
	raw, err := readZipFile(y)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(raw, &s); err != nil {
		return err
	}

	return nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
