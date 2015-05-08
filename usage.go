package juicage

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
)

type arg struct {
	name string
	help string
}

type Usage struct {
	flagSet  *flag.FlagSet
	name     string
	mandArgs []*arg
	optArgs  []*arg
	restArg  *arg
}

var CommandLineUsage = NewUsage(path.Base(os.Args[0]), flag.CommandLine)

func init() {
	flag.CommandLine.Usage = CommandLineUsage.Usage
}

func Mandatory(name string, help string) *Usage {
	return CommandLineUsage.Mandatory(name, help)
}

func Optional(name string, help string) *Usage {
	return CommandLineUsage.Optional(name, help)
}

func Rest(name string, help string) *Usage {
	return CommandLineUsage.Rest(name, help)
}

func Print() error {
	return CommandLineUsage.Print()
}

func String() string {
	return CommandLineUsage.String()
}

func NewUsage(name string, flagSet *flag.FlagSet) *Usage {
	ret := &Usage{flagSet, name, []*arg{}, []*arg{}, nil}
	flagSet.Usage = ret.Usage
	return ret
}

func (u *Usage) Mandatory(name string, help string) *Usage {
	u.mandArgs = append(u.mandArgs, &arg{name, help})
	return u
}

func (u *Usage) Optional(name string, help string) *Usage {
	u.optArgs = append(u.optArgs, &arg{name, help})
	return u
}

func (u *Usage) Rest(name string, help string) *Usage {
	u.restArg = &arg{name, help}
	return u
}

func (u *Usage) String() string {
	buf := &bytes.Buffer{}
	var longestName int

	fmt.Fprintf(buf, "usage: %s", u.name)
	for _, v := range u.mandArgs {
		fmt.Fprintf(buf, " %s", v.name)
		if len(v.name) > longestName {
			longestName = len(v.name)
		}
	}
	for _, v := range u.optArgs {
		fmt.Fprintf(buf, " [%s]", v.name)
		if len(v.name) > longestName {
			longestName = len(v.name)
		}
	}
	if u.restArg != nil {
		fmt.Fprintf(buf, " [%s]*", u.restArg.name)
		if len(u.restArg.name) > longestName {
			longestName = len(u.restArg.name)
		}
	}
	buf.WriteByte('\n')
	buf.WriteByte('\n')

	for _, v := range u.mandArgs {
		fmt.Fprintf(buf, " %*s: %s\n", longestName, v.name, v.help)
	}
	for _, v := range u.optArgs {
		fmt.Fprintf(buf, " %*s: %s\n", longestName, v.name, v.help)
	}
	if u.restArg != nil {
		fmt.Fprintf(buf, " %*s: %s\n", longestName, u.restArg.name, u.restArg.help)
	}

	buf.WriteByte('\n')
	//fmt.Fprintf()
	u.flagSet.SetOutput(buf)
	u.flagSet.PrintDefaults()
	return buf.String()
}

func (u *Usage) Print() error {
	_, err := fmt.Printf(u.String())
	if err != nil {
		return err
	}
	return nil
}

func (u *Usage) Usage() {
	u.Print()
	os.Exit(1)
}
