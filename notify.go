package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/guelfey/go.dbus"
)

type Urgency byte

const (
	LowUrgency Urgency = iota
	NormalUrgency
	CriticalUrgency
)

// notifications are displayed for the duration of timeout.
var timeout = 2 * time.Second

func Timeout() time.Duration       { return timeout }
func SetTimeout(dur time.Duration) { timeout = dur }

func Notify(name, summary, body string) error {
	return NotifyIcon(name, summary, body, "")
}

func NotifyUrgency(name, summary, body, icon string, urgency Urgency) error {
	hints := map[string]dbus.Variant{"urgency": dbus.MakeVariant(byte(urgency))}
	_, err := notify(name, summary, body, icon, 0, nil, hints, int32(timeout/time.Millisecond))
	return err
}

func NotifyIcon(name, summary, body, icon string) error {
	_, err := notify(name, summary, body, icon, 0, nil, nil, int32(timeout/time.Millisecond))
	return err
}

func notify(name, summary, body, icon string, replacesID uint32, actions []string, hints map[string]dbus.Variant, timeout int32) (id uint32, err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return 0, err
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, name, replacesID, icon, summary, body, actions, hints, timeout)
	if call.Err != nil {
		return 0, call.Err
	} else if call.Store(&id) != nil {
		return 0, errors.New("unrecognized response from notify daemon")
	}
	return
}

func main() {
	err := Notify("mywire", "Connection OK", "")
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
	NotifyUrgency("mywire", "Connection BAD", "", "", CriticalUrgency)
}
