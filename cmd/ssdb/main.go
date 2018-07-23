package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wzshiming/ssdb/term"
)

var auth = flag.String("p", "", "Authenticate password")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
Usage of %s:
	ssdb [options] [address]
	ssdb -p password 127.0.0.1:8888
`, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	err := term.Run(args[0], *auth)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
		return
	}
}
