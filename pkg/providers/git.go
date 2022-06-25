package providers

import (
	"context"
	"os"
	"path/filepath"

	"github.com/andersnormal/picasso/pkg/providers/iface"

	gg "github.com/go-git/go-git/v5"
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

	_, err = gg.PlainCloneContext(ctx, path, true, &gg.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	// r, err := gg.CloneContext(ctx, memory.NewStorage(), nil, &gg.CloneOptions{
	// 	Depth:    1,
	// 	URL:      url,
	// 	Progress: os.Stdout,
	// })
	// if err != nil {
	// 	return err
	// }

	// hash, err := r.ResolveRevision(plumbing.Revision("HEAD"))
	// if err != nil {
	// 	return err
	// }

	// // commit, err := r.CommitObject(*hash)
	// // if err != nil {
	// // 	return err
	// // }

	// // tree, err := commit.Tree()
	// // if err != nil {
	// // 	return err
	// // }

	// w, err := r.Worktree()
	// if err != nil {
	// 	return err
	// }

	// w.Filesystem.Chroot(path)
	// w.Checkout(&gg.CheckoutOptions{
	// 	Hash: *hash,
	// })

	// files := tree.Files()
	// err = files.ForEach(func(f *object.File) error {
	// 	parts := strings.Split(f.Name, string(os.PathSeparator))

	// 	// Store filename/path for returning and using later on
	// 	fpath := filepath.Join(path, filepath.Join(parts...))

	// 	fmt.Println(fpath, parts)

	// 	// if f.FileInfo().IsDir() {
	// 	// 	// Make Folder
	// 	// 	_ = os.MkdirAll(fpath, os.ModePerm)
	// 	// 	continue
	// 	// }

	// 	// // Make File
	// 	// if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
	// 	// 	return err
	// 	// }

	// 	// outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	// rc, err := f.Open()
	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	// _, err = io.Copy(outFile, rc)

	// 	// // Close the file without defer to close before next iteration of loop
	// 	// outFile.Close()
	// 	// rc.Close()

	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	// if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
	// 	// 	return err
	// 	// }

	// 	// file, err := os.Create(fpath)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	// fmt.Println(file)

	// 	b, err := f.IsBinary()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if !b {
	// 		fmt.Println(f.Name)
	// 	}

	// 	return nil
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}
