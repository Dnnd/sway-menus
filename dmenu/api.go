package dmenu

type MenuData interface {
	TotalLines() int
	AsString() string
}

type MenuPresenter interface {
	Show(data MenuData) (int, error)
}

type MenuFactory interface {
	NewMenuPresenter() MenuPresenter
}
