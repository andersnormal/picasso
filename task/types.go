package task

type Tasks map[string]*Task

type Task struct {
	Cmds        []*Cmd
	Deps        []*Dep
	Desc        string
	Dir         string
	Env         Vars
	Generates   []string
	IgnoreError bool `yaml:"ignore_error"`
	Method      string
	Prefix      string
	Silent      bool
	Sources     []string
	Status      []string
	Name        string
	Vars        Vars
}

type Cmd struct {
	Cmd         string
	Silent      bool
	Task        string
	Vars        Vars
	IgnoreError bool
}

type Vars map[string]Var

type Var struct {
	Static string
	Sh     string
}

type Dep struct {
	Task string
	Vars Vars
}
