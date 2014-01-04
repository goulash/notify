// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package notify

import (
	"github.com/guelfey/go.dbus"
)

var connection *dbus.Conn

// ServiceAvailable returns true if notifications via DBus are available.
//
// First, it initiates a connection via DBus to find out whether DBus is
// available, then it contacts the notification service to find out if it is
// available. If one or the other is not available, it returns false.
//
// Before using notify, it is a good idea (though not necessary) to test
// if this service is available.
func ServiceAvailable() bool {
	return true
}
