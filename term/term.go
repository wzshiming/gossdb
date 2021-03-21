package term

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	prompt "github.com/c-bata/go-prompt"
	// _ "github.com/wzshiming/winseq"
)

type CmdFunc func(cmd ...string) (string, error)

// Terminal Is a terminal renderer.
type Terminal struct {
	Reader  io.Reader
	Writer  io.Writer
	Prompt  string
	CmdFunc CmdFunc
}

// NewTerminal Create a new Terminal.
func NewTerminal(prompt string, cmd CmdFunc) *Terminal {
	return &Terminal{
		Reader:  os.Stdin,
		Writer:  os.Stdout,
		Prompt:  prompt,
		CmdFunc: cmd,
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.FindStartOfPreviousWord() == 0 {
		return prompt.FilterHasPrefix(command, d.GetWordBeforeCursorWithSpace(), true)
	}
	return []prompt.Suggest{}
}

// Run Is run the terminal.
func (c *Terminal) Run() error {
	logger := log.New(c.Writer, "", log.LstdFlags)
	pro := prompt.New(func(string) {}, completer,
		prompt.OptionPrefix(c.Prompt),
		prompt.OptionPrefixTextColor(prompt.DefaultColor),
		prompt.OptionPrefixBackgroundColor(prompt.DefaultColor),
	)

	for {
		line := pro.Input()
		read := csv.NewReader(bytes.NewBufferString(strings.TrimSpace(line)))
		read.Comma = ' '
		read.LazyQuotes = true
		read.TrimLeadingSpace = true
		da, err := read.ReadAll()
		if err != nil {
			logger.Println(err)
			continue
		}
		for _, v := range da {
			beg := time.Now()
			result, err := c.CmdFunc(v...)
			if err != nil {
				logger.Println(err)
				continue
			}
			sub := time.Since(beg).Truncate(time.Millisecond)
			fmt.Fprintln(c.Writer, result)
			fmt.Fprintf(c.Writer, "(%s)\n", sub)
		}
	}
	return nil
}
