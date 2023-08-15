package packages

import (
	"flag"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestCliWithArgs(t *testing.T) {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	ptr(argsWithoutProg)
	ptr(argsWithProg)
	ptr(arg)
}

func TestCliWithFlag(t *testing.T) {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	ptr("word:", *wordPtr)
	ptr("numb:", *numbPtr)
	ptr("fork:", *forkPtr)
	ptr("svar:", svar)
	ptr("tail:", flag.Args())
}

func TestCliSubCommand(t *testing.T) {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		ptr("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		ptr("subcommand 'foo'")
		ptr("  enable:", *fooEnable)
		ptr("  name:", *fooName)
		ptr("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		ptr("subcommand 'bar'")
		ptr("  level:", *barLevel)
		ptr("  tail:", barCmd.Args())
	default:
		ptr("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	ptr("FOO:", os.Getenv("FOO"))
	ptr("BAR:", os.Getenv("BAR"))

	ptr()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		ptr(pair[0])
	}
}

func TestCmd(t *testing.T) {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	ptr("> date")
	ptr(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			ptr("failed executing:", err)
		case *exec.ExitError:
			ptr("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	ptr("> grep hello")
	ptr(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	ptr("> ls -a -l -h")
	ptr(string(lsOut))
}

func TestExecProcess(t *testing.T) {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func TestSignal(t *testing.T) {
	worker := func() {
		ptr("working")
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			worker()
			time.Sleep(time.Second)
		}
	}()

	ptr("awaiting signal")
	ptr("exiting.", <-sigs)
}
