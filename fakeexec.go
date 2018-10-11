package fakeexec

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"testing"
)

var (
	// EnvName is used for
	EnvName = "FAKE_EXEC_PROCESS"

	TestFuncName = "TestFakeExecHandler"
)

func Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(os.Args[0], buildArgs(command, args)...)
	setupCmd(cmd)
	return cmd
}

func CommandContext(ctx context.Context, command string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, os.Args[0], buildArgs(command, args)...)
	setupCmd(cmd)
	return cmd
}

func buildArgs(command string, args []string) []string {
	return append([]string{"-test.run=" + TestFuncName, "--", command}, args...)
}

func setupCmd(cmd *exec.Cmd) {
	cmd.Env = append(cmd.Env, EnvName+"=1")
}

type HandleFunc func(cmd string, args []string) int

func Handle(t *testing.T, h HandleFunc) {
	if os.Getenv(EnvName) != "1" {
		t.Log("THIS IS NOT FAKE_EXEC PROCESS")
		return
	}

	args := flag.Args()
	code := h(args[0], args[1:])
	os.Exit(code)
}
