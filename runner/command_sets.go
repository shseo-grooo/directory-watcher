package runner

type Cmd string

func (c Cmd) String() string {
	return string(c)
}
type Path string

func (p Path) String() string {
	return string(p)
}

type CommandSet struct {
	Cmd Cmd
	Path Path
}

type CommandSets []CommandSet
