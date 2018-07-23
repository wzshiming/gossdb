package term

import (
	"fmt"
	"strings"

	"github.com/wzshiming/ssdb"
	ffmt "gopkg.in/ffmt.v1"
)

func Run(addr string, auth string) error {
	cli, err := ssdb.Connect(
		ssdb.Addr(addr),
		ssdb.Auth(auth),
	)
	if err != nil {
		return err
	}
	info, err := cli.Info()
	if err != nil {
		return err
	}
	fmt.Printf(welcome, info["version"])

	conn, err := Conn(cli)
	if err != nil {
		return err
	}

	return NewTerminal(fmt.Sprintf("\nSSDB %s > ", addr), NewExtra(conn).Cmd).Run()
}

func Conn(cli *ssdb.Client) (CmdFunc, error) {
	return func(cmd ...string) (string, int, error) {
		if len(cmd) == 0 {
			return "", 0, nil
		}
		ss := ssdb.Values{}
		for _, list := range cmd {
			val, err := ssdb.NewValue(list)
			if err != nil {
				return "", 0, err
			}
			ss = append(ss, val)
		}
		val, err := cli.Do(ss)
		if err != nil {
			return "", 0, err
		}
		val, err = ssdb.ResultProcessing(val, err)
		if err != nil {
			return err.Error(), 0, nil
		}
		if val == nil {
			return "not found", 0, nil
		}
		key := strings.Replace(strings.ToLower(cmd[0]), "_", "", -1)
		if key == "info" {
			return val[1:].String(), 0, nil
		}

		if kind[key] != "map" {
			return val.String(), len(val), nil
		}
		table := [][]string{
			{"key", "value"},
			{"---", "-----"},
		}
		for i := 0; i < len(val); i += 2 {
			table = append(table, []string{val[i].String(), val[i+1].String()})
		}
		return strings.Join(ffmt.FmtTable(table), "\n"), len(val) / 2, nil

	}, nil
}
