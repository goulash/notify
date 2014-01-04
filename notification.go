// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package notify

import (
	"errors"
	"time"

	"github.com/guelfey/go.dbus"
)

var connection *dbus.Conn

type NotificationUrgency byte

const (
	LowUrgency NotificationUrgency = iota
	NormalUrgency
	CriticalUrgency
)

func (u NotificationUrgency) asHint() map[string]dbus.Variant {
	return map[string]dbus.Variant{"urgency": dbus.MakeVariant(byte(u))}
}

type Notification struct {
	Name    string
	Summary string
	Body    string

	IconPath string
	Timeout  time.Duration
	Urgency  NotificationUrgency
}

func New(name, summary, body, icon string, timeout time.Duration, urgency NotificationUrgency) *Notification {
	return &Notification{name, summary, body, icon, timeout, urgency}
}

func (n Notification) Send() (id uint32, err error) {
	return notify(n.Name, n.Summary, n.Body, n.IconPath, 0, nil, n.Urgency.asHint(), n.TimeoutInMS())
}

func (n Notification) SendMsg(summary, body string) (id uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, 0, nil, n.Urgency.asHint(), n.TimeoutInMS())
}

func (n Notification) SendUrgentMsg(summary, body string, urgency NotificationUrgency) (id uint32, err error) {
	return notify(n.Name, summary, body, n.IconPath, 0, nil, urgency.asHint(), n.TimeoutInMS())
}

func (n Notification) Replace(id uint32) error {
	_, err := notify(n.Name, n.Summary, n.Body, n.IconPath, id, nil, n.Urgency.asHint(), n.TimeoutInMS())
	return err
}

func (n Notification) ReplaceMsg(id uint32, summary, body string) error {
	_, err := notify(n.Name, summary, body, n.IconPath, id, nil, n.Urgency.asHint(), n.TimeoutInMS())
	return err
}

func (n Notification) ReplaceUrgentMsg(id uint32, summary, body string, urgency NotificationUrgency) error {
	_, err := notify(n.Name, summary, body, n.IconPath, id, nil, urgency.asHint(), n.TimeoutInMS())
	return err
}

func (n Notification) TimeoutInMS() int32 {
	return int32(n.Timeout / time.Millisecond)
}

func notify(name, summary, body, icon string, replacesID uint32, actions []string, hints map[string]dbus.Variant, timeout int32) (id uint32, err error) {
	if connection == nil {
		connection, err := dbus.SessionBus()
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
