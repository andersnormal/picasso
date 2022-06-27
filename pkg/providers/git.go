package providers

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/andersnormal/picasso/pkg/providers/iface"
	"github.com/andersnormal/picasso/pkg/specs"
	"github.com/andersnormal/picasso/pkg/tmpl"
	"gopkg.in/yaml.v2"

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

	var s *specs.Tmpl
	t := tmpl.New()

	// Find spec ...
	if err := ff.ForEach(func(f *object.File) error {
		if !strings.Contains(f.Name, specs.TmplFile) {
			return nil
		}

		r, err := f.Reader()
		if err != nil {
			return err
		}

		err = yaml.NewDecoder(r).Decode(&s)
		if err != nil {
			return err
		}

		err = t.ApplyPrompts(s.Inputs)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	// :-) just resetting the iterator
	ff.Close()
	ff, err = ref.Files()
	if err != nil {
		return err
	}

	// Find spec ...
	if err := ff.ForEach(func(f *object.File) error {
		parts := strings.Split(f.Name, string(os.PathSeparator))
		fpath := filepath.Join(path, filepath.Join(parts...))

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

		ok, err := f.IsBinary()
		if err != nil {
			return err
		}

		var ignore bool
		for _, p := range s.Ignores {
			if strings.Contains(f.Name, p) {
				ignore = true
				break
			}
		}

		r, err := f.Reader()
		if err != nil {
			return err
		}

		if !ignore && !ok {
			text, err := f.Contents()
			if err != nil {
				return err
			}

			out, err := t.Apply(text)
			if err != nil {
				return err
			}

			r := strings.NewReader(out)

			_, err = io.Copy(outFile, r)
			if err != nil {
				return err
			}
		} else {
			_, err = io.Copy(outFile, r)
			if err != nil {
				return err
			}
		}

		return err
	}); err != nil {
		return err
	}

	return nil
}
