package swaymsg

const Executable = "swaymsg"

type Swaymsg interface {
	Send() ([]byte, error)
}
