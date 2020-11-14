package swaymsg

type Swaymsg interface {
	Send() (string, error)
}

