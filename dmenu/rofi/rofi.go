package rofi

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"io"
	"os/exec"
	"strconv"
)

type ErrRofiFailed struct {
	Output string
	Cause  error
}

type ShowMenuWithRofi struct{}

type SwitchWorkspacesMenuFactory struct{}

const Executable = "rofi"

func (f SwitchWorkspacesMenuFactory) NewMenu() dmenu.MenuPresenter {
	return &ShowMenuWithRofi{}
}

func (e *ErrRofiFailed) Error() string {
	return fmt.Sprintf("rofi failed: %s", e.Output)
}

func (w ShowMenuWithRofi) Show(data dmenu.MenuData) (string, error) {
	cmd := exec.Command(Executable, "-dmenu", "-l", strconv.Itoa(data.Lines), "-i", "-format", "i")
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
		return "", &ErrRofiFailed{Output: string(out), Cause: err}
	}
	return string(out), nil
}
