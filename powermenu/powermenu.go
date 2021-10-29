package powermenu

import (
	"errors"
	"github.com/Dnnd/sway-menus/dmenu"
	generic "github.com/Dnnd/sway-menus/domain/generic_menu"
	"github.com/Dnnd/sway-menus/logind"
)

type PowermenuSerivce struct {
	menu     generic.GenericMenu
	logind   *logind.Manager
	launcher dmenu.MenuPresenter
}

const PowerOffTag generic.EntryTag = "poweroff"
const RebootTag generic.EntryTag = "reboot"
const HibernateTag generic.EntryTag = "hibernate"
const Suspend generic.EntryTag = "suspend"
const Logout generic.EntryTag = "logout"

func NewPowerMenuService(menu generic.GenericMenu, logind *logind.Manager, launcher dmenu.MenuPresenter) *PowermenuSerivce {
	return &PowermenuSerivce{menu: menu, logind: logind, launcher: launcher}
}

func (p *PowermenuSerivce) Run() error {
	matchedEntry, err := p.launcher.Show(p.menu)
	if err != nil {
		return err
	}
	var matched = p.menu[matchedEntry]
	switch matched.Tag {
	case HibernateTag:
		return p.logind.Hibernate()
	case Logout:
		return p.logind.TerminateCurrentSession()
	case Suspend:
		return p.logind.Suspend()
	case RebootTag:
		return p.logind.Reboot()
	case PowerOffTag:
		return p.logind.PowerOff()
	}
	return errors.New("unknown menu tag")
}
