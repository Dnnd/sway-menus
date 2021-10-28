package wofi

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"io"
	"os/exec"
	"strconv"
)

type ErrWofiFailed struct {
	Output string
	Cause  error
}

type ShowMenuWithWofi struct{}

type SwitchWorkspacesMenuFactory struct{}

const Executable = "wofi"

func (f SwitchWorkspacesMenuFactory) NewMenu() dmenu.MenuPresenter {
	return &ShowMenuWithWofi{}
}

func (e *ErrWofiFailed) Error() string {
	return fmt.Sprintf("wofi failed: %s", e.Output)
}

func (w ShowMenuWithWofi) Show(data dmenu.MenuData) (string, error) {
	cmd := exec.Command(Executable, "--dmenu", "-L", strconv.Itoa(data.Lines))
	in, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	go func() {
		defer in.Close()
		io.WriteString(in, data.Entries)
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &ErrWofiFailed{Output: string(out), Cause: err}
	}
	return string(out), nil
}
