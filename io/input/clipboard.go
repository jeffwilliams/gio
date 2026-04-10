// SPDX-License-Identifier: Unlicense OR MIT

package input

import (
	"io"
	"slices"

	"gioui.org/io/clipboard"
	"gioui.org/io/event"
)

// clipboardState contains the state for clipboard event routing.
type clipboardState struct {
	receivers []event.Tag
}

type clipboardQueue struct {
	// request avoid read clipboard every frame while waiting.
	requested bool
	mime      string
	text      []byte
}

// WriteClipboard returns the most recent data to be copied
// to the clipboard, if any.
func (q *clipboardQueue) WriteClipboard() (mime string, content []byte, ok bool) {
	if q.text == nil {
		return "", nil, false
	}
	content = q.text
	q.text = nil
	return q.mime, content, true
}

// ClipboardRequested reports if any new handler is waiting
// to read the clipboard.
func (q *clipboardQueue) ClipboardRequested(state clipboardState) bool {
	req := len(state.receivers) > 0 && q.requested
	log("gio: clipboardQueue.ClipboardRequested: on return q.requested = false and len(state.receivers) = %d\n", len(state.receivers))
	log("gio: stacktrace: %s\n", stack(false))
	return req
}

func (q *clipboardQueue) SetClipboardNotRequested() {
	q.requested = false
}

func (q *clipboardQueue) Push(state clipboardState, e event.Event) (clipboardState, []taggedEvent) {
	var evts []taggedEvent
	for _, r := range state.receivers {
		evts = append(evts, taggedEvent{tag: r, event: e})
	}
	state.receivers = nil
	log("gio: clipboardQueue.Push: on return len(state.receivers) = 0\n")
	return state, evts
}

func (q *clipboardQueue) ProcessWriteClipboard(req clipboard.WriteCmd) {
	defer req.Data.Close()
	content, err := io.ReadAll(req.Data)
	if err != nil {
		return
	}
	q.mime = req.Type
	q.text = content
}

func (q *clipboardQueue) ProcessReadClipboard(state clipboardState, tag event.Tag) clipboardState {

	if !q.requested && slices.Contains(state.receivers, tag) {
		log("gio: issue reproduced. ProcessReadClipboard was called when q.requested is false but state.receivers contains the tag\n")
		log("gio: stacktrace: %s\n", stack(false))
		IssueReproduced()
	}
	
	// Disable original fix; testing the fix for when ReadClipboard returns false
	//q.requested = true
	if slices.Contains(state.receivers, tag) {
		log("gio: clipboardQueue.ProcessReadClipboard: state.receivers contains tag. Returning same state\n")
		return state
	}
	n := len(state.receivers)
	state.receivers = append(state.receivers[:n:n], tag)
	// Original fix comments this line out:
	q.requested = true
	log("gio: clipboardQueue.ProcessReadClipboard: state.receivers updated to contain %d items\n", len(state.receivers))
	return state
}
