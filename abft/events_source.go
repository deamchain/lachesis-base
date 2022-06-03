package abft

import (
	"github.com/deamchain/lachesis-base/hash"
	"github.com/deamchain/lachesis-base/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
