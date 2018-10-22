package monitor

import (
	"fmt"
	"log"
	"path/filepath"
	"swarm/source"
	"time"

	"github.com/rjeczalik/notify"
)

// FilterFn is the shape of a function that can used as a filter for a Monitor
type FilterFn func(notify.Event, string) bool

// Monitor is used to recursively watch for file changes within a workspace
type Monitor struct {
	workspace *source.Workspace
	channel   chan notify.EventInfo
	filter    FilterFn
}

// NewMonitor creates a new Monitor
func NewMonitor(workspace *source.Workspace, filter FilterFn) *Monitor {
	channel := make(chan notify.EventInfo, 2048)

	rootPathRecursive := filepath.Join(workspace.RootPath(), "./...")
	if err := notify.Watch(rootPathRecursive, channel, (notify.Write | notify.Remove)); err != nil {
		log.Fatal(err)
	}

	return &Monitor{
		workspace,
		channel,
		filter,
	}
}

const notifyInterval = 10 * time.Minute
const debounceInterval = 100 * time.Millisecond

// NotifyOnChanges notifies when events occur (after debouncing)
func (mon *Monitor) NotifyOnChanges(callback func(changes *EventChangeset)) {
	debounceTimer := time.NewTimer(notifyInterval)
	changeset := NewEventChangeset()

	var e notify.EventInfo
	var start time.Time
	for {
		select {
		case e = <-mon.channel:
			// receive an event
			event := e.Event()
			path := e.Path()
			if mon.filter == nil || mon.filter(event, path) {
				if changeset.empty() {
					start = time.Now()
					fmt.Print("Change detected...")
				} else {
					fmt.Print(".")
				}
				changeset.Add(event, path)
				debounceTimer.Reset(debounceInterval)
			}

		case <-debounceTimer.C:
			// debounce and fire callback
			if changeset.nonEmpty() {
				fmt.Println("")
				callback(changeset)
				elapsed := time.Since(start)
				fmt.Printf("...done in %s\n", elapsed)
			} else {
				fmt.Println("")
				fmt.Println("...no changes")
			}
			changeset = NewEventChangeset()
		}
	}
}

// Stop cancels the recursive watcher
func (mon *Monitor) Stop() {
	notify.Stop(mon.channel)
}
