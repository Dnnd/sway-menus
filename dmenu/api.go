package dmenu

type DmenuCommand interface {
	Show() (string, error)
}
