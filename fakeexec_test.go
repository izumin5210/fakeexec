package fakeexec_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/izumin5210/fakeexec"
)

func TestFakeExec(t *testing.T) {
	out, _ := fakeexec.Command("echo", "foo").Output()

	if got, want := string(out), "bar\n"; got != want {
		t.Errorf("printed %q, want %q", got, want)
	}
}

func TestFakeExecHandler(t *testing.T) {
	fakeexec.Handle(t, func(cmd string, _ []string) int {
		if cmd == "echo" {
			fmt.Fprintln(os.Stdout, "bar")
		}
		return 0
	})
}
