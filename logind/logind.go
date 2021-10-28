package logind

import (
	"github.com/godbus/dbus/v5"
	"os"
)

type Manager struct {
	conn *dbus.Conn
	obj  dbus.BusObject
}

const LogindNode = "org.freedesktop.login1"
const LogindPath = "/org/freedesktop/login1"
const LogindInterface = "org.freedesktop.login1.Manager"
const TerminateSession = LogindInterface + ".TerminateSession"
const Reboot = LogindInterface + ".Reboot"
const Suspend = LogindInterface + ".Suspend"
const Hibernate = LogindInterface + ".Hibernate"
const PowerOff = LogindInterface + ".PowerOff"

func NewLogindManager(conn *dbus.Conn) Manager {
	managerObj := conn.Object(LogindNode, LogindPath)
	return Manager{conn: conn, obj: managerObj}
}

func ConnectToLogindManager() (Manager, error) {
	conn, err := dbus.SystemBus()
	return NewLogindManager(conn), err
}

func (m Manager) TerminateSession(sessionID string) error {
	call := m.obj.Call(TerminateSession, 0, sessionID)
	return call.Err
}

func (m Manager) Reboot() error {
	return m.obj.Call(Reboot, 0).Err
}

func (m Manager) PowerOff() error {
	return m.obj.Call(PowerOff, 0).Err
}

func (m Manager) Suspend() error {
	return m.obj.Call(Suspend, 0).Err
}

func (m Manager) Hibernate() error {
	return m.obj.Call(Hibernate, 0).Err
}

func getCurrentSession() string {
	return os.Getenv("XDG_SESSION_ID")
}

func (m Manager) TerminateCurrentSession() error {
	currentSessionID := getCurrentSession()
	return m.TerminateSession(currentSessionID)
}
