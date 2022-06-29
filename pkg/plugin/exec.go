package plugin

import "context"

// var stdin bytes.Buffer
// 		var stdout bytes.Buffer
// 		spec := &specs.Tmpl{Version: "0.0.2"}

// 		path, err := os.Getwd()
// 		if err != nil {
// 			return err
// 		}

// 		// plugins ...
// 		c := exec.Command(filepath.Join(path, "examples", "gen-tmpl", "main"))
// 		c.Stdin = &stdin
// 		c.Stdout = &stdout

// 		quit := make(chan struct{})

// 		go func() {
// 			err := c.Run()
// 			if err != nil {
// 				l.Error("failed to run plugin", zap.Error(err))
// 			}

// 			fmt.Println("gere")

// 			quit <- struct{}{}
// 		}()

// 		enc := gob.NewEncoder(&stdin)

// 		err = enc.Encode(spec)
// 		if err != nil {
// 			return err
// 		}

// 		<-quit

// 		fmt.Println(stdout.String())

// Executor ...
type Executor interface {
	ExecWithContext(context.Context) error
}

type executor struct{}

// NewExecutor ...
func NewExecutor() Executor {
	return &executor{}
}

// ExecWithContext ...
func (e *executor) ExecWithContext(ctx context.Context) error {
	return nil
}
