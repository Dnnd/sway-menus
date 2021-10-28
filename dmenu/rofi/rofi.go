package rofi

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type ErrRofiFailed struct {
	Output string
	Cause  error
}

type ShowMenuWithRofi struct{}

type MenuFactory struct{}

const Executable = "rofi"

func (f MenuFactory) NewMenuPresenter() dmenu.MenuPresenter {
	return &ShowMenuWithRofi{}
}

func (e *ErrRofiFailed) Error() string {
	return fmt.Sprintf("rofi failed: %s", e.Output)
}

func (w ShowMenuWithRofi) Show(data dmenu.MenuData) (int, error) {
	cmd := exec.Command(Executable, "-dmenu", "-l", strconv.Itoa(data.TotalLines()), "-i", "-format", "i", "-no-custom")
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
		return 0, &ErrRofiFailed{Output: string(out), Cause: err}
	}
	entryNum, err := strconv.Atoi(strings.Trim(string(out), "\n"))
	return entryNum, nil
}
