// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package notify

import (
	"errors"

	"github.com/godbus/dbus"
)

// connection is a global D-Bus connection. TODO: I do not know if this
// can be used concurrently!
var connection *dbus.Conn

// ServiceAvailable returns true if notifications via DBus are available.
//
// First, it initiates a connection via DBus to find out whether DBus is
// available, then it contacts the notification service to find out if it is
// available. If one or the other is not available, it returns false.
//
// Before using notify, it is a good idea (though not necessary) to test
// if this service is available. If it's not available, this does not
// tell you why though. Maybe another day.
func ServiceAvailable() bool {
	if connection == nil {
		var err error
		connection, err = dbus.SessionBus()
		if err != nil {
			return false
		}
	}

	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.GetCapabilities", 0)
	if call.Err != nil {
		return false
	}

	return true
}

// notify does the real work of getting a connection and talking to the
// notification daemon. It doesn't really talk though.
//
// To have some elements use their defaults, the following is accepted:
//
//	name = ""
//	body = ""
//	replacesID = 0
//	actions = nil
//	hints = nil
//
// So you see, really only summary and timeout are required for a meaningful
// notification.
func notify(name, summary, body, icon string, replacesID uint32, actions []string, hints map[string]dbus.Variant, timeout int32) (id uint32, err error) {
	if connection == nil {
		connection, err = dbus.SessionBus()
		if err != nil {
			return 0, err
		}
	}

	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, name, replacesID, icon, summary, body, actions, hints, timeout)
	if call.Err != nil {
		return 0, call.Err
	} else if call.Store(&id) != nil {
		return 0, errors.New("unrecognized response from notify daemon")
	}
	return
}
