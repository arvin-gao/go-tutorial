package packages

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestCliWithArgs(t *testing.T) {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	pln(argsWithoutProg)
	pln(argsWithProg)
	pln(arg)
}

func TestCliWithFlag(t *testing.T) {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	pln("word:", *wordPtr)
	pln("numb:", *numbPtr)
	pln("fork:", *forkPtr)
	pln("svar:", svar)
	pln("tail:", flag.Args())
}

func TestCliSubCommand(t *testing.T) {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		pln("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		pln("subcommand 'foo'")
		pln("  enable:", *fooEnable)
		pln("  name:", *fooName)
		pln("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		pln("subcommand 'bar'")
		pln("  level:", *barLevel)
		pln("  tail:", barCmd.Args())
	default:
		pln("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	pln("FOO:", os.Getenv("FOO"))
	pln("BAR:", os.Getenv("BAR"))

	pln()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		pln(pair[0])
	}
}
