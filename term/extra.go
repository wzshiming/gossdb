package term

import (
	"os"
	"strings"
)

type Extra struct {
	Map   map[string]CmdFunc
	Other CmdFunc
}

func (e *Extra) Cmd(cmd ...string) (string, error) {
	if len(cmd) == 0 {
		return "", nil
	}
	fun := e.Map[strings.ToLower(cmd[0])]
	if fun != nil {
		return fun(cmd...)
	}
	return e.Other(cmd...)
}

func (e *Extra) AddCmd(name string, fun CmdFunc) {
	e.Map[strings.ToLower(name)] = fun
}

func NewExtra(other CmdFunc) *Extra {
	e := &Extra{
		Map:   map[string]CmdFunc{},
		Other: other,
	}
	e.AddCmd("quit", quit)
	e.AddCmd("help", help)
	return e
}

func quit(cmd ...string) (string, error) {
	os.Exit(0)
	return "", nil
}

func help(cmd ...string) (string, error) {
	return usage, nil
}
