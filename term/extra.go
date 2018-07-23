package term

import (
	"os"
	"strings"
)

type Extra struct {
	Map   map[string]CmdFunc
	Other CmdFunc
}

func (e *Extra) Cmd(cmd ...string) (string, int, error) {
	if len(cmd) == 0 {
		return "", 0, nil
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

func quit(cmd ...string) (string, int, error) {
	os.Exit(0)
	return "", 0, nil
}

func help(cmd ...string) (string, int, error) {
	return usage, 0, nil
}
