package swaymsg

import (
	"fmt"
	"os/exec"
)

type ErrFocusWindowFailed struct {
	Output string
	Cause  error
}

func (e *ErrFocusWindowFailed) Error() string {
	return fmt.Sprintf("swaymsg focus failed: %s, %s", e.Cause, e.Output)
}

type FocusWindow struct {
	Id int
}

func NewFocusWindowMessage(id int) Swaymsg {
	return &FocusWindow{
		Id: id,
	}
}

func (f *FocusWindow) Send() ([]byte, error) {
	focusChosenWindow := exec.Command(Executable, fmt.Sprintf("[con_id=%d]", f.Id), "focus")
	out, err := focusChosenWindow.CombinedOutput()
	if err != nil {
		return nil, &ErrFocusWindowFailed{
			Output: string(out),
			Cause:  err,
		}
	}
	return out, nil
}
