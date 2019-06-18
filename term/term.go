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
	s := []prompt.Suggest{
		{Text: "quit", Description: "quit ssdb"},
		{Text: "help", Description: "help ssdb"},

		{Text: "auth", Description: "auth password"},
		{Text: "dbsize", Description: "dbsize"},
		{Text: "flushdb", Description: "flushdb [type]"},
		{Text: "info", Description: "info [opt]"},
		{Text: "slaveof", Description: "slaveof id host port [auth last_seq last_key]"},

		{Text: "list_allow_ip", Description: "list_allow_ip ip_rule"},
		{Text: "add_allow_ip", Description: "add_allow_ip ip_rule"},
		{Text: "del_allow_ip", Description: "del_allow_ip ip_rule"},
		{Text: "list_deny_ip", Description: "list_deny_ip ip_rule"},
		{Text: "add_deny_ip", Description: "add_deny_ip ip_rule"},
		{Text: "del_deny_ip", Description: "del_deny_ip ip_rule"},

		{Text: "set", Description: "set key value"},
		{Text: "setx", Description: "setx key value ttl"},
		{Text: "setnx", Description: "setnx key value"},
		{Text: "expire", Description: "expire key ttl"},
		{Text: "ttl", Description: "ttl key"},
		{Text: "get", Description: "get key"},
		{Text: "getset", Description: "getset key value"},
		{Text: "del", Description: "del key"},
		{Text: "incr", Description: "incr key [num]"},
		{Text: "exists", Description: "exists key"},
		{Text: "getbit", Description: "getbit key offset"},
		{Text: "setbit", Description: "setbit key offset val"},
		{Text: "bitcount", Description: "bitcount key [start] [end]"},
		{Text: "countbit", Description: "countbit key start size"},
		{Text: "substr", Description: "substr key start size"},
		{Text: "strlen", Description: "strlen key"},
		{Text: "keys", Description: "keys key_start key_end limit"},
		{Text: "rkeys", Description: "rkeys key_start key_end limit"},
		{Text: "scan", Description: "scan key_start key_end limit"},
		{Text: "rscan", Description: "rscan key_start key_end limit"},
		{Text: "multi_set", Description: "multi_set key1 value1 key2 value2 ..."},
		{Text: "multi_get", Description: "multi_get key1 key2 ..."},
		{Text: "multi_del", Description: "multi_del key1 key2 ..."},

		{Text: "hset", Description: "hset name key value"},
		{Text: "hget", Description: "hget name key"},
		{Text: "hdel", Description: "hdel name key"},
		{Text: "hincr", Description: "hincr name key [num]"},
		{Text: "hexists", Description: "hexists name key"},
		{Text: "hsize", Description: "hsize name"},
		{Text: "hlist", Description: "hlist name_start name_end limit"},
		{Text: "hrlist", Description: "hrlist name_start name_end limit"},
		{Text: "hkeys", Description: "hkeys name key_start key_end"},
		{Text: "hgetall", Description: "hgetall name"},
		{Text: "hscan", Description: "hscan name key_start key_end limit"},
		{Text: "hrscan", Description: "hrscan name key_start key_end limit"},
		{Text: "hclear", Description: "hclear name"},
		{Text: "multi_hset", Description: "multi_hset name key1 value1 key2 value2 ..."},
		{Text: "multi_hget", Description: "multi_hget name key1 key2 ..."},
		{Text: "multi_hdel", Description: "multi_hdel name key1 key2 ..."},

		{Text: "zset", Description: "zset name key score"},
		{Text: "zget", Description: "zget name key"},
		{Text: "zdel", Description: "zdel name key"},
		{Text: "zincr", Description: "zincr name key num"},
		{Text: "zexists", Description: "zexists name key"},
		{Text: "zsize", Description: "zsize name"},
		{Text: "zlist", Description: "zlist name_start name_end limit"},
		{Text: "zrlist", Description: "zrlist name_start name_end limit"},
		{Text: "zkeys", Description: "zkeys name key_start score_start score_end limit"},
		{Text: "zscan", Description: "zscan name key_start score_start score_end limit"},
		{Text: "zrscan", Description: "zrscan name key_start score_start score_end limit"},
		{Text: "zrank", Description: "zrank name key"},
		{Text: "zrrank", Description: "zrrank name key"},
		{Text: "zrange", Description: "zrange name offset limit"},
		{Text: "zrrange", Description: "zrrange name offset limit"},
		{Text: "zclear", Description: "zclear name"},
		{Text: "zcount", Description: "zcount name score_start score_end"},
		{Text: "zsum", Description: "zsum name score_start score_end"},
		{Text: "zavg", Description: "zavg name score_start score_end"},
		{Text: "zremrangebyrank", Description: "zremrangebyrank name start end"},
		{Text: "zremrangebyscore", Description: "zremrangebyscore name start end"},
		{Text: "zpop_front", Description: "zpop_front name limit"},
		{Text: "zpop_back", Description: "zpop_back name limit"},
		{Text: "multi_zset", Description: "multi_zset name key1 score1 key2 score2 ..."},
		{Text: "multi_zget", Description: "multi_zget name key1 key2 ..."},
		{Text: "multi_zdel", Description: "multi_zdel name key1 key2 ..."},

		{Text: "qpush_front", Description: "qpush_front name item1 item2 ..."},
		{Text: "qpush_back", Description: "qpush_back name item1 item2 ..."},
		{Text: "qpop_front", Description: "qpop_front name size"},
		{Text: "qpop_back", Description: "qpop_back name size"},
		{Text: "qpush", Description: "qpush name item1 item2 ..."},
		{Text: "qpop", Description: "qpop name size"},
		{Text: "qfront", Description: "qfront name"},
		{Text: "qback", Description: "qback name"},
		{Text: "qsize", Description: "qsize name"},
		{Text: "qclear", Description: "qclear name"},
		{Text: "qget", Description: "qget name index"},
		{Text: "qset", Description: "qset name index val"},
		{Text: "qrange", Description: "qrange name offset limit"},
		{Text: "qslice", Description: "qslice name begin end"},
		{Text: "qtrim_front", Description: "qtrim_front name size"},
		{Text: "qtrim_back", Description: "qtrim_back name size"},
		{Text: "qlist", Description: "qlist name_start name_end limit"},
		{Text: "qrlist", Description: "qrlist name_start name_end limit"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursorWithSpace(), true)
}

// Run Is run the terminal.
func (c *Terminal) Run() error {
	fmt.Fprintln(c.Writer, welcome)
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
			sub := time.Now().Sub(beg).Truncate(time.Millisecond)
			fmt.Fprintln(c.Writer, result)
			fmt.Fprintf(c.Writer, "(%s)\n", sub)
		}
	}
	return nil
}
