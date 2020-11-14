package swaymsg

import (
	"fmt"
	"os/exec"
)

type ErrGetTreeFailed struct {
	Output string
	Cause  error
}

func (e *ErrGetTreeFailed) Error() string {
	return fmt.Sprintf("swaymsg -t get_tree failed: %s, %s", e.Cause, e.Output)
}

type GetTree struct {
}

func NewGetTree() Swaymsg {
	return &GetTree{
	}
}

func (f *GetTree) Send() ([]byte, error) {
	focusChosenWindow := exec.Command(Executable, "-t", "get_tree")
	out, err := focusChosenWindow.CombinedOutput()
	if err != nil {
		return nil, &ErrFocusWindowFailed{
			Output: string(out),
			Cause:  err,
		}
	}
	return out, nil
}
