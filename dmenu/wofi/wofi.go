package wofi

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type ErrWofiFailed struct {
	Output string
	Cause  error
}

type ShowMenuWithWofi struct{}

type MenuFactory struct{}

const Executable = "wofi"

func (f MenuFactory) NewMenuPresenter() dmenu.MenuPresenter {
	return &ShowMenuWithWofi{}
}

func (e *ErrWofiFailed) Error() string {
	return fmt.Sprintf("wofi failed: %s", e.Output)
}

func (w ShowMenuWithWofi) Show(data dmenu.MenuData) (int, error) {
	cmd := exec.Command(Executable, "--dmenu", "-L", strconv.Itoa(data.TotalLines()))
	in, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	go func() {
		defer in.Close()
		io.WriteString(in, data.AsString())
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, &ErrWofiFailed{Output: string(out), Cause: err}
	}
	entryNum, err := strconv.Atoi(strings.Trim(string(out), "\n"))
	return entryNum, nil
}
