# fakeexec

Wrapper of [Testing `os/exec.Command`](https://npf.io/2015/06/testing-exec-command/)

## Example

```go
var execCommand = exec.Command

func EchoFoo() ([]byte, error) {
	return ExecCommand("echo", "foo).Output()
}
```

```go
func TestEcho(t *testing.T) {
	execCommand = fakeexec.Command
	defer func() { execCommand = exec.Command }()

	out, _ := EchoFoo()

	if got, want := string(out), "foo"; got != want {
		t.Errorf("prints %q, want %q", got, want)
	}
}

// You should define "TestFakeExecHandler" to handle executed commands.
func TestFakeExecHandler(t *testing.T) {
	fakeexec.Handle(t, func(cmd string, args []string) int {
		if cmd == "echo" {
			fmt.Fprintln(os.Stdout, args[len(args)-1])
		}
		return 0
	})
}
```
