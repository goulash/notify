// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package notify

import (
	"time"

	"github.com/godbus/dbus"
)

// NotificationUrgency can be either LowUrgency, NormalUrgency, and CriticalUrgency.
// It is conceivable that some notification daemons make no distinction between the
// different urgencies, but enough do that it makes sense to use them.
type NotificationUrgency byte

const (
	LowUrgency      NotificationUrgency = iota // LowUrgency probably shouldn't even be shown ;-)
	NormalUrgency                              // NormalUrgency is for information that is interesting.
	CriticalUrgency                            // CriticalUrgency is for errors or severe events.
)

// asHint returns the NotificationUrgency in the type that the DBus
// specification requires.
func (u NotificationUrgency) asHint() map[string]dbus.Variant {
	return map[string]dbus.Variant{"urgency": dbus.MakeVariant(byte(u))}
}

// Notification is there to provide you with full power of your notifications.
// It is possible for you to use a Notification as you use the notify library
// without them. This allows for multiple defaults.
//
// For example:
//
//	func main() {
//		critical := notify.New("prog", "", "", "critical-icon.png", time.Duration(0), notify.CriticalUrgency)
//		boring := notify.New("prog", "", "", "low-icon.png", 1 * time.Second, notify.LowUrgency)
//		boring.SendMsg("Nothing is happening... boring!", "")
//		critical.SendMsg("Your computer is on fire!", "Here is what you should do:\n ...")
//	}
//
type Notification struct {
	// Name represents the application name sending the notification.  This is
	// optional and can be the empty string "".
	Name string
	// Summary represents the subject of the notification.
	Summary string
	// Body represents the main body with extra details. Some notification
	// daemons ignore the body; it is optional and can be the empty string "".
	Body string

	// IconPath is a path to an icon that should be used for the notification.
	// Some notification daemons ignore the icon path; it is optional and can
	// be the empty string "".
	IconPath string
	// Timeout is the requested timeout for the notification. Some notification
	// daemons override the requested timeout. A value of 0 is a request that
	// it not timeout at all.
	Timeout time.Duration
	// Urgency determines the urgency of the notification, which can be one of
	// LowUrgency, NormalUrgency, and CriticalUrgency.
	Urgency NotificationUrgency
}

// New returns a pointer to a new Notification.
func New(name, summary, body, icon string, timeout time.Duration, urgency NotificationUrgency) *Notification {
	return &Notification{name, summary, body, icon, timeout, urgency}
}

// Send sends the notification n as it is, and returns the ID and err, possibly
// nil.
func (n Notification) Send() (id uint32, err error) {
	return notify(n.Name, n.Summary, n.Body, n.IconPath, 0, nil, n.Urgency.asHint(), n.timeoutInMS())
}

// SendMsg is identical to notify.SendMsg, except that the rest of the values
// come from n.
func (n Notification) SendMsg(summary, body string) (id uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, 0, nil, n.Urgency.asHint(), n.timeoutInMS())
}

// SendUrgentMsg is identical to notify.SendUrgentMsg, except that the rest of
// the values come from n.
func (n Notification) SendUrgentMsg(summary, body string, urgency NotificationUrgency) (id uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, 0, nil, urgency.asHint(), n.timeoutInMS())
}

// Replace replaces the notification with the ID id as it is, and returns the
// new ID and an error if one occured.
func (n Notification) Replace(id uint32) (newID uint32, err error) {
	return notify(n.Name, n.Summary, n.Body, n.IconPath, id, nil, n.Urgency.asHint(), n.timeoutInMS())
}

// ReplaceMsg is identical to notify.ReplaceMsg, except that the rest of the
// values come from n.
func (n Notification) ReplaceMsg(id uint32, summary, body string) (newID uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, id, nil, n.Urgency.asHint(), n.timeoutInMS())
}

// ReplaceUrgentMsg is identical to notify.ReplaceUrgentMsg, except that the
// rest of the values come from n.
func (n Notification) ReplaceUrgentMsg(id uint32, summary, body string, urgency NotificationUrgency) (newID uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, id, nil, urgency.asHint(), n.timeoutInMS())
}

// timeoutInMS returns Timeout in milliseconds.
//
// The specification specifies that the timeout is the number of milliseconds
// that the notification should be displayed.
func (n Notification) timeoutInMS() int32 {
	return int32(n.Timeout / time.Millisecond)
}
