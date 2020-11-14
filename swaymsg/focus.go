package swaymsg

import (
	"fmt"
	"os/exec"
)

const SwaymsgExecutable = "swaymsg"

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

func (f *FocusWindow) Send() (string, error) {
	focusChosenWindow := exec.Command(SwaymsgExecutable, fmt.Sprintf("[con_id=%d]", f.Id), "focus")
	out, err := focusChosenWindow.CombinedOutput()
	if err != nil {
		return "", &ErrFocusWindowFailed{
			Output: string(out),
			Cause:  err,
		}
	}
	return string(out), nil
}
