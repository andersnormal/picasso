package plugin

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"google.golang.org/protobuf/proto"
)

type executor struct{}

// Executor ...
type Executor interface {
	// ExecWithContext ...
	ExecWithContext(context.Context, string, *PluginRequest) error
}

// NewExecutor ...
func NewExecutor() Executor {
	return &executor{}
}

// ExecWithContext ...
func (e *executor) ExecWithContext(ctx context.Context, p string, req *PluginRequest) error {
	exec := exec.CommandContext(ctx, p)
	stdin, err := exec.StdinPipe()
	if err != nil {
		return err
	}

	m, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewReader(m))
	}()

	bb, err := exec.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(bb))

	msg := &PluginResponse{}
	err = proto.Unmarshal(bb, msg)
	if err != nil {
		return err
	}

	if msg.Error != "" {
		return fmt.Errorf(msg.Error)
	}

	return nil
}
