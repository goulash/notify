// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package notify provides tray notifications via D-Bus freedesktop.org notifications.
//
// WARNING: You are free to use this package, however please note that the API is
// likely to be unstable till all the functionality is implemented, and there is
// a lot missing!
//
// There is not much you need to do use this package, just start sending
// notifications:
//
//	func main() {
//		id, err := notify.SendMsg("My summary", "My body, which some notification services don't show")
//		if err != nil {
//			fmt.Println(err)
//		}
//		// You can also ignore the error message, in which case failed notifications are ignored,
//		// because let's be honest, who cares if the notifications get delivered or not? ;-)
//		notify.ReplaceUrgentMsg("Forget that <i>last message</i>", "", notify.LowUrgency)
//	}
//
// The defaults are a timeout of 3 seconds for notifications and a normal
// urgency, with everything being empty.  You are free to change these via
// Init, SetName, SetTimeout, SetIconPath, and SetUrgency.
//
// Alternatively, you can create your own Notification template, via New.
//
// The notify package has been developed according to
// https://developer.gnome.org/notification-spec, although there is a lot of
// functionality missing.
//
// There is some markup that you can use in the summary and body parts of a
// notification:
//
//	<b>bold</b>
//	<i>italic</i>
//	<u>underline</u>
//	<a href="...">hyperlink</a>
//	<img src="..." alt="..." />
//
// I have tried to implement this in the Go philosophy, please let me know if
// I can improve it somehow.
package notify

import (
	"time"
)

// note acts as the default notification, which allows you to set default
// parameters and then send messages without creating any Notifications.
var note = Notification{
	Timeout: 3 * time.Second,
	Urgency: NormalUrgency,
}

// Init sets the defaults for the implicit notification.
func Init(name, icon string, timeout time.Duration, urgency NotificationUrgency) {
	note.Name = name
	note.IconPath = icon
	note.Timeout = timeout
	note.Urgency = urgency
}

// Name returns the name for the implicit notification.
func Name() string { return note.Name }

// SetName sets the name for the implicit notification.
func SetName(name string) { note.Name = name }

// IconPath returns the icon path for the implicit notification.
func IconPath() string { return note.IconPath }

// SetIconPath sets the icon path for the implicit notification.
func SetIconPath(path string) { note.IconPath = path }

// Timeout returns the timeout for the implicit notification.
func Timeout() time.Duration { return note.Timeout }

// SetTimeout sets the timeout for the implicit notification.
func SetTimeout(dur time.Duration) { note.Timeout = dur }

// Urgency returns the urgency level for the implicit notification.
// It can be either LowUrgency, NormalUrgency, or CriticalUrgency.
func Urgency() NotificationUrgency { return note.Urgency }

// SetUrgency sets the urgency level for the implicit notification.
// It can be either LowUrgency, NormalUrgency, or CriticalUrgency.
func SetUrgency(urgency NotificationUrgency) { note.Urgency = urgency }

// SendMsg sends the summary and the body as a notification, returning a unique
// notification ID and an error, possibly nil. It takes all other values from
// the implicit notification object.
func SendMsg(summary, body string) (id uint32, err error) {
	return note.SendMsg(summary, body)
}

// SendUrgentMsg sends the summary and the body as a notification with the
// urgency of urgency, and returns a unique notification ID and an error,
// possibly nil. Otherwise it is like SendMsg.
func SendUrgentMsg(summary, body string, urgency NotificationUrgency) (id uint32, err error) {
	return note.SendUrgentMsg(summary, body, urgency)
}

// ReplaceMsg replaces the already existing notification with the ID id with
// summary and body, returning an error if it fails. It takes all other values
// from the implicit notification object.
//
// In particular, if the notification it is replacing had other properties,
// such as another urgency, these are also replaced by the defaults in the
// implicit notification!
func ReplaceMsg(id uint32, summary, body string) error {
	return note.ReplaceMsg(id, summary, body)
}

// ReplaceUrgentMsg replaces the already existing notification with the ID id
// with summary and body and urgency, returning an error if it fails. It takes
// all other values from the implicit notification object.
func ReplaceUrgentMsg(id uint32, summary, body string, urgency NotificationUrgency) error {
	return note.ReplaceUrgentMsg(id, summary, body, urgency)
}
