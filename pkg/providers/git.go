package providers

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/andersnormal/picasso/pkg/providers/iface"

	gg "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"go.uber.org/zap"
)

type git struct {
	opts *iface.ProviderOpts
}

// NewGit ...
func NewGit(opts ...iface.ProviderOpt) iface.Provider {
	options := new(iface.ProviderOpts)

	g := new(git)
	g.opts = options
	g.opts.Configure(opts...)

	return g
}

// CloneWithContext ...
func (g *git) CloneWithContext(ctx context.Context, url string, folder string) error {
	ll := g.opts.Logger.With(zap.String("provider", "git"), zap.String("url", url))

	ll.Info("Cloning repository")

	path, err := filepath.Abs(folder)
	if err != nil {
		return err
	}

	r, err := gg.CloneContext(ctx, memory.NewStorage(), nil, &gg.CloneOptions{
		URL:   url,
		Depth: 1,
	})
	if err != nil {
		return err
	}

	head, err := r.Head()
	if err != nil {
		return err
	}

	ref, err := r.CommitObject(head.Hash())
	if err != nil {
		return err
	}

	ff, err := ref.Files()
	if err != nil {
		return err
	}

	if err := ff.ForEach(func(f *object.File) error {
		parts := strings.Split(f.Name, string(os.PathSeparator))
		fpath := filepath.Join(path, filepath.Join(parts...))

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		mode, err := f.Mode.ToOSFileMode()
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
		if err != nil {
			return err
		}
		defer outFile.Close()

		r, err := f.Reader()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, r)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
