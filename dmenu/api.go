package dmenu

type MenuData struct {
	Entries string
	Lines   int
}

type MenuPresenter interface {
	Show(data MenuData) (string, error)
}

type MenuFactory interface {
	NewMenu() MenuPresenter
}
