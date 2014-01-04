// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package notify

import (
	"github.com/goulash/notify"
	"time"
)

// This is a simple example for how to use the notify package.
func Example() {
	notify.SetName("Simple")

	notify.SendMsg("Starting up the Simple Server", "")
	time.Sleep(3 * time.Second)
	id, _ = notify.SendUrgentMsg("Oops, made a big mistake!", "", notify.CriticalUrgency)
	time.Sleep(1 * time.Second)
	notify.ReplaceMsg(id, "Ha! Fixed that, thank goodness!", "")
}
