// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package notify provides DBus notifications.
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

func Init(name, icon string, timeout time.Duration, urgency NotificationUrgency) {
	note.Name = name
	note.IconPath = icon
	note.Timeout = timeout
	note.Urgency = urgency
}

func Name() string        { return note.Name }
func SetName(name string) { note.Name = name }

func IconPath() string        { return note.IconPath }
func SetIconPath(path string) { note.IconPath = path }

func Timeout() time.Duration       { return note.Timeout }
func SetTimeout(dur time.Duration) { note.Timeout = dur }

func Urgency() NotificationUrgency           { return note.Urgency }
func SetUrgency(urgency NotificationUrgency) { note.Urgency = urgency }

func SendMsg(summary, body string) (id uint32, err error) {
	return note.SendMsg(summary, body)
}

func SendUrgentMsg(summary, body string, urgency NotificationUrgency) (id uint32, err error) {
	return note.SendUrgentMsg(summary, body, urgency)
}

func ReplaceMsg(id uint32, summary, body string) error {
	return note.ReplaceMsg(id, summary, body)
}

func ReplaceUrgentMsg(id uint32, summary, body string, urgency NotificationUrgency) error {
	return note.ReplaceUrgentMsg(id, summary, body, urgency)
}
