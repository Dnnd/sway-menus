package wofi

import (
	"fmt"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"github.com/Dnnd/sway-window-switcher/domain/workspace"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type WofiDmenu struct {
	Menu workspace.Dmenu
}

type ErrWofiFailed struct {
	Output string
	Cause  error
}

const WOFI_EXECUTABLE = "wofi"

func NewWofiDmenuCommand(menu workspace.Dmenu) dmenu.DmenuCommand {
	return &WofiDmenu{Menu: menu}
}

func (e *ErrWofiFailed) Error() string {
	return fmt.Sprintf("wofi failed: %s", e.Output)
}

func (w *WofiDmenu) Show() (string, error) {
	cmd := exec.Command(WOFI_EXECUTABLE, "--dmenu", "-L", strconv.Itoa(w.Menu.Lines))
	in, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	go func() {
		defer in.Close()
		io.WriteString(in, w.Menu.Entries)
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &ErrWofiFailed{Output: string(out), Cause: err}
	}
	return string(out), nil
}

func ParseWofiDmenuEntryNumber(dmenuOutput string) (int, error) {
	num, err := strconv.ParseInt(strings.Trim(dmenuOutput, "\n"), 10, 32)
	return int(num), err
}
